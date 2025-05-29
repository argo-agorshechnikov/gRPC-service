package main

import (
	"context"
	"log"
	"time"

	pb "github.com/argo-agorshechnikov/gRPC-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Set connect with grpc server
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewExampleServiceClient(conn)

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call method GetData
	response, err := client.GetData(ctx, &pb.Request{
		Text: "Test message for server!",
	})
	if err != nil {
		log.Fatalf("Error call GetData: %v", err)
	}

	log.Printf("Server response: %s", response.Result)

}
