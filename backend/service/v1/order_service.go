package v1

import (
	v1pb "buf-protoc/proto/gen/go/v1"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// OrderService implements the order service.
type OrderService struct {
	v1pb.UnimplementedOrderServiceServer
}

// NewOrderService creates a new OrderService.
func NewOrderService() *OrderService {
	return &OrderService{}
}

// 创建订单
func (s *OrderService) CreateOrder(ctx context.Context, in *v1pb.CreateOrderRequest) (*v1pb.Order, error) {
	fmt.Printf("CreateOrder %+v", in)

	return &v1pb.Order{
		Id:         "1",
		No:         "202101010001",
		CreateTime: timestamppb.Now(),
		UpdateTime: timestamppb.Now(),
		Status:     v1pb.OrderStatus_ORDER_STATUS_CREATED,
		Amount:     100.34,
	}, nil
}

// 获取订单
func (s *OrderService) GetOrder(ctx context.Context, in *v1pb.GetOrderRequest) (*v1pb.Order, error) {
	fmt.Printf("GetOrder %+v", in)

	return &v1pb.Order{
		Id:         "1",
		No:         "202101010001",
		CreateTime: timestamppb.Now(),
		UpdateTime: timestamppb.Now(),
		Status:     v1pb.OrderStatus_ORDER_STATUS_CREATED,
		Amount:     100.34,
	}, nil
}

// 删除订单
func (s *OrderService) DeleteOrder(ctx context.Context, in *v1pb.DeleteOrderRequest) (*v1pb.OperateResult, error) {
	fmt.Printf("DeleteOrder %+v", in)

	return &v1pb.OperateResult{
		Code:    0,
		Message: "success",
		Details: []string{},
	}, nil
}
