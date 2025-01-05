package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/bufbuild/protovalidate-go"

	v1pb "buf-protoc/proto/gen/go/v1"

	"google.golang.org/genproto/googleapis/api/annotations"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	grpcruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/soheilhy/cmux"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	v1pb.UnimplementedHelloServiceServer

	echoServer *echo.Echo
	grpcServer *grpc.Server
	muxServer  cmux.CMux

	grpcMux  *grpcruntime.ServeMux
	grpcConn *grpc.ClientConn

	port int

	cancel context.CancelFunc
}

func NewServer(port int) *Server {
	server := &Server{}
	if port < 80 {
		port = 36789
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		panic(err)
	}

	m := cmux.New(lis)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(authInterceptor),
		// Override the maximum receiving message size to 100M for uploading large sheets.
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.InitialWindowSize(100000000),
		grpc.InitialConnWindowSize(100000000),
	)

	// 要获取proto文件中的服务描述符，可以通过以下方式：
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Register the gRPC server.
	v1pb.RegisterHelloServiceServer(grpcServer, server)

	// Create Echo instance.
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create HTTP server using grpc-gateway.
	ctx := context.Background()

	//  gatewayModifier := auth.GatewayResponseModifier{Store: s.store}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	mux := grpcruntime.NewServeMux(
		grpcruntime.WithRoutingErrorHandler(func(ctx context.Context, sm *grpcruntime.ServeMux, m grpcruntime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
			if httpStatus != http.StatusNotFound {
				grpcruntime.DefaultRoutingErrorHandler(ctx, sm, m, w, r, httpStatus)
				return
			}

			err := &grpcruntime.HTTPStatusError{
				HTTPStatus: httpStatus,
				Err:        status.Errorf(codes.NotFound, "Routing error. Please check the request URI %v", r.RequestURI),
			}

			grpcruntime.DefaultHTTPErrorHandler(ctx, sm, m, w, r, err)
		}),
	)

	err = v1pb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", port), opts)
	if err != nil {
		fmt.Printf("failed to register handler: %v\n", err)
		panic(err)
	}

	// REST gateway proxy.
	grpcEndpoint := fmt.Sprintf(":%d", port)
	grpcConn, err := grpc.Dial(
		grpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(100*1024*1024), // Set MaxCallRecvMsgSize to 100M so that users can receive up to 100M via REST calls.
		),
	)
	if err != nil {
		panic(err)
	}

	// Register grpc-gateway mux with Echo
	e.Any("/*", echo.WrapHandler(mux))

	return &Server{
		echoServer: e,
		grpcServer: grpcServer,
		muxServer:  m,
		grpcMux:    mux,
		grpcConn:   grpcConn,
	}
}

func (s *Server) Run() {
	_, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	// Match gRPC connections.
	// grpcL := s.muxServer.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))

	grpcL := s.muxServer.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// Match HTTP connections.
	httpL := s.muxServer.Match(cmux.HTTP1Fast())

	go func() {

		fmt.Println("starting gRPC server")
		if err := s.grpcServer.Serve(grpcL); err != nil {
			panic(err)
		}
	}()

	go func() {
		fmt.Println("starting HTTP server")
		httpServer := &http.Server{
			Handler: s.echoServer,
		}
		if err := httpServer.Serve(httpL); err != nil {
			panic(err)
		}
	}()

	// Start cmux serving.
	fmt.Printf("starting cmux server")
	if err := s.muxServer.Serve(); err != nil {
		panic(err)
	}
}

// Shutdown will shut down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("Stopping web server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Cancel the worker
	if s.cancel != nil {
		s.cancel()
	}

	// Shutdown echo
	if s.grpcServer != nil {
		stopped := make(chan struct{})
		go func() {
			s.grpcServer.GracefulStop()
			close(stopped)
		}()

		t := time.NewTimer(1 * time.Second)
		select {
		case <-t.C:
			s.grpcServer.Stop()
		case <-stopped:
			t.Stop()
		}
	}
	if s.echoServer != nil {
		if err := s.echoServer.Shutdown(ctx); err != nil {
			s.echoServer.Logger.Fatal(err)
		}
	}
	if s.muxServer != nil {
		s.muxServer.Close()
	}

	return nil
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	for k, v := range md {
		slog.Info("metadata", k, v)
	}
	// ... 验证 token 的逻辑 ...
	// if !isValidToken(token) {
	// 	return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	// }
	// 验证通过，继续调用后续处理函数
	return handler(ctx, req)
}

func (s *Server) GetUser(ctx context.Context, req *v1pb.Req) (*v1pb.User, error) {
	// 从 metadata 中获取方法名
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	for k, v := range md {
		slog.Info("metadata", k, v)
	}

	if err := protovalidate.Validate(req); err != nil {
		fmt.Println("error:\n", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &v1pb.User{
		Name: "John Doe",
	}, nil
}

// printMethodOptions 打印方法描述符中的所有选项
func printMethodOptions(method protoreflect.MethodDescriptor) {
	// 获取方法选项
	options := method.Options().(*descriptorpb.MethodOptions)

	// 打印 google.api.http 选项
	if proto.HasExtension(options, annotations.E_Http) {
		httpRule := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
		fmt.Printf("google.api.http: %v\n", httpRule)
	}

	// 打印 google.api.method_signature 选项
	if proto.HasExtension(options, annotations.E_MethodSignature) {
		methodSignature := proto.GetExtension(options, annotations.E_MethodSignature).([]string)
		fmt.Printf("google.api.method_signature: %v\n", methodSignature)
	}

	// 打印 bytebase.v1.permission 选项
	if proto.HasExtension(options, v1pb.E_Permission) {
		permission := proto.GetExtension(options, v1pb.E_Permission).(string)
		fmt.Printf("beats.v1.permission: %v\n", permission)
	}

	// 打印 bytebase.v1.auth_method 选项
	if proto.HasExtension(options, v1pb.E_AuthMethod) {
		authMethod := proto.GetExtension(options, v1pb.E_AuthMethod).(v1pb.AuthMethod)
		fmt.Printf("beats.v1.auth_method: %v\n", authMethod)
	}
}
