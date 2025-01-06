package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"buf-protoc/api/auth"
	"buf-protoc/component/state"

	"buf-protoc/component/config"

	"github.com/bufbuild/protovalidate-go"

	v1pb "buf-protoc/proto/gen/go/v1"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	grpcruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/reflect/protoregistry"

	"github.com/soheilhy/cmux"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	v1pb.UnimplementedHelloServiceServer

	echoServer *echo.Echo
	grpcServer *grpc.Server
	muxServer  cmux.CMux

	grpcMux *grpcruntime.ServeMux

	port int

	cancel context.CancelFunc
}

func NewServer(port int, profile *config.Profile) *Server {
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {

		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			service := services.Get(i)
			if serviceHandler, _ := proto.GetExtension(service.Options(), v1pb.E_ServiceExtend).(*v1pb.ServiceExtend); serviceHandler != nil {
				fmt.Println("servicename:" + string(service.FullName()))
				fmt.Println("serviceSignature:" + serviceHandler.ServiceSignature)
				fmt.Println("permission:" + serviceHandler.Permission)
				fmt.Println("authmethod:" + serviceHandler.AuthMethod.String())
				fmt.Println("audit:" + strconv.FormatBool(serviceHandler.Audit))
			}

			methods := service.Methods()
			for k := 0; k < methods.Len(); k++ {
				method := methods.Get(k)
				if methodHandler, _ := proto.GetExtension(method.Options(), v1pb.E_MethodExtend).(*v1pb.MethodExtend); methodHandler != nil {
					fmt.Println("methodname:" + string(method.Name()))
					fmt.Println("methodSignature:" + methodHandler.MethodSignature)
					fmt.Println("permission:" + methodHandler.Permission)
					fmt.Println("authmethod:" + methodHandler.AuthMethod.String())
					fmt.Println("audit:" + strconv.FormatBool(methodHandler.Audit))
				}
			}

		}

		return true
	})

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

	state, err := state.New()
	if err != nil {
		slog.Error("failed to create state", "error", err)
		panic(err)
	}

	authProvider := auth.New("", state, profile)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(authProvider.AuthenticationInterceptor),
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

	// Register grpc-gateway mux with Echo
	e.Any("/*", echo.WrapHandler(mux))

	server.echoServer = e
	server.grpcServer = grpcServer
	server.muxServer = m
	server.grpcMux = mux
	server.port = port

	server.configHttpMiddleware(ctx)

	return server
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

func (s *Server) GetUser(ctx context.Context, req *v1pb.Req) (*v1pb.User, error) {
	if err := protovalidate.Validate(req); err != nil {
		fmt.Println("error:\n", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &v1pb.User{
		Name: "John Doe",
	}, nil
}
