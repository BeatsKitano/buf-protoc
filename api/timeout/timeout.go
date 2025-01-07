package timeout

import (
	v1pb "buf-protoc/proto/gen/go/v1"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type overrideStream struct {
	childCtx context.Context
	grpc.ServerStream
}

func (s overrideStream) Context() context.Context {
	return s.childCtx
}

// APIAuthInterceptor is the rate_limit and timeout interceptor for gRPC server.
type APITimeoutInterceptor struct {
	methoder map[string]*v1pb.MethodExtend
}

// New creates a new APITimeoutInterceptor.
func New(methoder map[string]*v1pb.MethodExtend) *APITimeoutInterceptor {
	return &APITimeoutInterceptor{
		methoder: methoder,
	}
}

// TimeoutInterceptor is the unary interceptor for gRPC API.
func (in *APITimeoutInterceptor) UnaryServerInterceptor(ctx context.Context, request any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	fullName := serverInfo.FullMethod

	extend, ok := in.methoder[fullName]
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, "method %s is not found rate limit exceeded", fullName)
	}

	if extend.Timeout <= 0 {
		return handler(ctx, request)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(extend.Timeout)*time.Millisecond)
	defer cancel()

	// 创建一个 channel 用于接收处理结果
	resChan := make(chan interface{})
	errChan := make(chan error)
	go func() {
		res, err := handler(timeoutCtx, request)
		if err != nil {
			errChan <- err
		} else {
			resChan <- res
		}
	}()

	select {
	case <-timeoutCtx.Done():
		return nil, status.Errorf(codes.DeadlineExceeded, "request timed out")
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return nil, err
	}
}

func (in *APITimeoutInterceptor) UnaryServerStreamInterceptor(request any, ss grpc.ServerStream, serverInfo *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	fullName := serverInfo.FullMethod

	extend, ok := in.methoder[fullName]
	if !ok {
		return status.Errorf(codes.ResourceExhausted, "method %s is not found rate limit exceeded", fullName)
	}

	sss := overrideStream{ServerStream: ss, childCtx: ctx}
	if extend.Timeout <= 0 {
		return handler(ctx, sss)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(extend.Timeout)*time.Millisecond)
	defer cancel()

	// 创建一个 channel 用于接收处理结果
	errChan := make(chan error)
	go func() {
		err := handler(timeoutCtx, sss)
		errChan <- err
	}()

	select {
	case <-timeoutCtx.Done():
		return status.Errorf(codes.DeadlineExceeded, "request timed out")
	case err := <-errChan:
		return err
	}
}
