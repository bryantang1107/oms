package main

import (
	"context"
	"log"
	"net"

	common "github.com/bryantang1107/commons"
	"google.golang.org/grpc"
)

var grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")

func main() {
	// gRPC server to receive request
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen to : $v", grpcAddr)
	}

	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGRPCHandler(grpcServer, service)

	service.CreateOrder(context.Background())

	log.Println("GRPC Server Started At : ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
