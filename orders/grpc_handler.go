package main

import (
	"context"
	"log"

	pb "github.com/bryantang1107/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	gHandler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, gHandler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order received : %v", payload)
	order := &pb.Order{
		ID: "42",
	}
	return order, nil
}
