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

	grpcruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
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
		// Override the maximum receiving message size to 100M for uploading large sheets.
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.InitialWindowSize(100000000),
		grpc.InitialConnWindowSize(100000000),
	)

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

func (s *Server) GetUser(ctx context.Context, req *v1pb.Req) (*v1pb.User, error) {
	fmt.Println("GetUser")
	if err := protovalidate.Validate(req); err != nil {
		fmt.Println("error:\n", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &v1pb.User{
		Name: "John Doe",
	}, nil
}
