package ratelimit

import (
	v1pb "buf-protoc/proto/gen/go/v1"
	"context"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sync"
)

type overrideStream struct {
	childCtx context.Context
	grpc.ServerStream
}

func (s overrideStream) Context() context.Context {
	return s.childCtx
}

// APIAuthInterceptor is the rate_limit and timeout interceptor for gRPC server.
type APIRateLimitInterceptor struct {
}

var limiterPool sync.Map

// New creates a new APIRateLimitInterceptor.
func New(methoder map[string]*v1pb.MethodExtend) *APIRateLimitInterceptor {
	for k, v := range methoder {
		if v.Rpm > 0 {
			limiter := rate.NewLimiter(rate.Limit(v.Rpm*1.0/60.0), int(v.Rpm))
			limiterPool.Store(k, limiter)
		}
	}

	return &APIRateLimitInterceptor{}
}

// RateLimitInterceptor is the unary interceptor for gRPC API.
func (in *APIRateLimitInterceptor) UnaryServerInterceptor(ctx context.Context, request any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	fullName := serverInfo.FullMethod

	limiterAny, ok := limiterPool.Load(fullName)
	if ok { // 如果存在限流器，则进行限流
		limiter := limiterAny.(*rate.Limiter)
		if !limiter.Allow() {
			return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
		}
	}

	return handler(ctx, request)
}

func (in *APIRateLimitInterceptor) UnaryServerStreamInterceptor(request any, ss grpc.ServerStream, serverInfo *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	fullName := serverInfo.FullMethod

	sss := overrideStream{ServerStream: ss, childCtx: ctx}

	limiterAny, ok := limiterPool.Load(fullName)
	if ok {
		limiter := limiterAny.(*rate.Limiter)
		if !limiter.Allow() {
			return status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
		}
		return handler(ctx, sss)
	}

	return handler(ctx, sss)

}
