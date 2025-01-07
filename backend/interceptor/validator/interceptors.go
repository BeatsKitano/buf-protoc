package validator

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// UnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := protovalidate.Validate(req.(proto.Message)); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that validates incoming messages.
//
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
// The stage at which invalid messages will be rejected with `InvalidArgument` varies based on the
// type of the RPC. For `ServerStream` (1:m) requests, it will happen before reaching any userspace
// handlers. For `ClientStream` (n:1) or `BidiStream` (n:m) RPCs, the messages will be rejected on
// calls to `stream.Recv()`.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{
			ServerStream: stream,
		}

		return handler(srv, wrapper)
	}
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m any) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}

	if err := protovalidate.Validate(m.(proto.Message)); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
