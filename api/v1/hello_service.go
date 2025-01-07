package v1

import (
	v1pb "buf-protoc/proto/gen/go/v1"
	"context"
)

// HelloService implements the actuator service.
type HelloService struct {
	v1pb.UnimplementedHelloServiceServer
}

// NewHelloService creates a new HelloService.
func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) GetUser(ctx context.Context, req *v1pb.Req) (*v1pb.User, error) {
	return &v1pb.User{
		Name: "John Doe",
	}, nil
}
