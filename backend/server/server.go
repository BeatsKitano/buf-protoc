package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"buf-protoc/backend/interceptor/auth"
	"buf-protoc/backend/interceptor/ratelimit"
	"buf-protoc/backend/interceptor/timeout"
	"buf-protoc/backend/interceptor/validator"

	"buf-protoc/component/state"

	"buf-protoc/component/config"

	v1pb "buf-protoc/proto/gen/go/v1"

	apiv1 "buf-protoc/api/v1"

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

	port string

	echoServer     *echo.Echo
	grpcServer     *grpc.Server
	grpcGatewayMux *grpcruntime.ServeMux
	muxServer      cmux.CMux

	cancel context.CancelFunc

	// 静态解析所有的proto文件，获取所有的服务和方法
	// key: "/"+service.FullName + "/"+ + method.Name
	methodExtends map[string]*v1pb.MethodExtend
}

func NewServer(port string, profile *config.Profile) *Server {
	server := &Server{
		methodExtends: make(map[string]*v1pb.MethodExtend, 0),
		port:          port,
	}

	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			service := services.Get(i)

			sn := service.FullName()
			methods := service.Methods()
			for k := 0; k < methods.Len(); k++ {
				method := methods.Get(k)
				if methodExtend, ok := proto.GetExtension(method.Options(), v1pb.E_MethodExtend).(*v1pb.MethodExtend); ok && methodExtend != nil {
					key := fmt.Sprintf("/%s/%s", sn, method.Name())
					server.methodExtends[key] = methodExtend
				}
			}
		}

		return true
	})

	for k, v := range server.methodExtends {
		fmt.Println(k, v)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		panic(err)
	}

	server.muxServer = cmux.New(lis)

	state, err := state.New()
	if err != nil {
		slog.Error("failed to create state", "error", err)
		panic(err)
	}

	authProvider := auth.New(server.methodExtends, "", state, profile)
	ratelimitProvider := ratelimit.New(server.methodExtends)
	timeoutProvider := timeout.New(server.methodExtends)

	grpc.EnableTracing = true
	server.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authProvider.UnaryServerInterceptor,
			ratelimitProvider.UnaryServerInterceptor,
			timeoutProvider.UnaryServerInterceptor,
			validator.UnaryServerInterceptor(),
		),
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.InitialWindowSize(100000000),
		grpc.InitialConnWindowSize(100000000),
	)

	// 要获取proto文件中的服务描述符，可以通过以下方式：
	// Register reflection service on gRPC server.
	reflection.Register(server.grpcServer)

	// Create Echo instance.
	server.echoServer = echo.New()
	server.echoServer.HideBanner = true
	server.echoServer.HidePort = false
	server.echoServer.Use(middleware.Logger())
	server.echoServer.Use(middleware.Recover())

	// Create HTTP server using grpc-gateway.
	server.grpcGatewayMux = grpcruntime.NewServeMux(
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

	ctx, cancel := context.WithCancel(context.Background())
	server.cancel = cancel

	err = v1pb.RegisterHelloServiceHandlerFromEndpoint(ctx, server.grpcGatewayMux, port, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		fmt.Printf("failed to register handler: %v\n", err)
		panic(err)
	}

	// Register the gRPC server.
	v1pb.RegisterHelloServiceServer(server.grpcServer, apiv1.NewHelloService())

	// Register grpc-gateway mux with Echo
	server.echoServer.Any("/*", echo.WrapHandler(server.grpcGatewayMux))

	return server
}

func (s *Server) Run() {
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
	fmt.Println("starting cmux server")
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
