package permission

import (
	"buf-protoc/backend/component/iam"
	v1pb "buf-protoc/proto/gen/go/v1"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PermissionInterceptor struct {
	methoder   map[string]*v1pb.MethodExtend
	iamManager *iam.Manager
}

func NewPermissionInterceptor(iamManager *iam.Manager, methoder map[string]*v1pb.MethodExtend) *PermissionInterceptor {
	return &PermissionInterceptor{
		iamManager: iamManager,
		methoder:   methoder,
	}
}

// UnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
func (i *PermissionInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		extend, ok := i.methoder[info.FullMethod]
		if !ok {
			return nil, status.Error(codes.Unimplemented, "method not found")
		}

		if extend.Permission == "" {
			return handler(ctx, req)
		}

		if ok, err := i.iamManager.CheckPermission(ctx, extend.Permission); err != nil {
			return nil, status.Error(codes.PermissionDenied, "permission denied")
		} else if !ok {
			return nil, status.Error(codes.PermissionDenied, "permission denied")
		}

		return handler(ctx, req)
	}
}
