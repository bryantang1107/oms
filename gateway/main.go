package main

import (
	"log"
	"net/http"

	common "github.com/bryantang1107/commons"
	pb "github.com/bryantang1107/commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// use gRPC to route to microservices (LB)
var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {
	// Initialize connection (without transport security) to order service
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: $v", err)
	}

	log.Println("Dialing order service at : ", orderServiceAddr)
	orderServiceClient := pb.NewOrderServiceClient(conn)

	// ensure connection is closed after running
	defer conn.Close()

	// Expose http server
	mux := http.NewServeMux()
	// handle order service
	handler := NewHandler(orderServiceClient)
	// handle incoming endpointa
	handler.registerRoutes(mux)

	log.Printf("Starting http server at %v", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start up server")
	}
}
