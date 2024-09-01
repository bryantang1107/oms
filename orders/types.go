package main

import (
	"context"

	pb "github.com/bryantang1107/commons/api"
)

type OrderService interface {
	CreateOrder(context context.Context) error
	ValidateOrder(context context.Context, p *pb.CreateOrderRequest) error
}

type OrderStore interface {
	Create(context context.Context) error
}
