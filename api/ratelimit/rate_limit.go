package ratelimit

import (
	v1pb "buf-protoc/proto/gen/go/v1"
	"context"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gvisor.dev/gvisor/pkg/sync"
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
	methoder map[string]*v1pb.MethodExtend

	limiter sync.Map
}

// New creates a new APIRateLimitInterceptor.
func New(methoder map[string]*v1pb.MethodExtend) *APIRateLimitInterceptor {
	return &APIRateLimitInterceptor{
		methoder: methoder,
	}
}

// RateLimitInterceptor is the unary interceptor for gRPC API.
func (in *APIRateLimitInterceptor) RateLimitInterceptor(ctx context.Context, request any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	fullName := serverInfo.FullMethod

	extend, ok := in.methoder[fullName]
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, "method %s is not found rate limit exceeded", fullName)
	}

	if extend.Rpm <= 0 {
		return handler(ctx, request)
	}

	limiterAny, ok := in.limiter.Load(fullName)
	if !ok {
		ratePerSecond := rate.Limit(extend.Rpm) / 60
		burst := extend.Rpm / 60
		limiter := rate.NewLimiter(ratePerSecond, int(burst))
		in.limiter.Store(fullName, limiter)
		return handler(ctx, request)
	}

	limiter := limiterAny.(*rate.Limiter)
	if !limiter.Allow() {
		return nil, status.Errorf(codes.ResourceExhausted, "method %s rate limit exceeded", fullName)
	}

	return handler(ctx, request)
}

func (in *APIRateLimitInterceptor) RateLimitStreamInterceptor(request any, ss grpc.ServerStream, serverInfo *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	fullName := serverInfo.FullMethod

	extend, ok := in.methoder[fullName]
	if !ok {
		return status.Errorf(codes.ResourceExhausted, "method %s is not found rate limit exceeded", fullName)
	}

	sss := overrideStream{ServerStream: ss, childCtx: ctx}

	if extend.Rpm <= 0 {
		return handler(ctx, sss)
	}

	limiterAny, ok := in.limiter.Load(fullName)
	if !ok {
		ratePerSecond := rate.Limit(extend.Rpm) / 60
		burst := extend.Rpm / 60
		limiter := rate.NewLimiter(ratePerSecond, int(burst))
		in.limiter.Store(fullName, limiter)
		return handler(ctx, sss)
	}

	limiter := limiterAny.(*rate.Limiter)
	if !limiter.Allow() {
		return status.Errorf(codes.ResourceExhausted, "method %s rate limit exceeded", fullName)
	}

	return handler(ctx, sss)

}
