package main

import (
	"context"
	"log"
	"net"

	pb "github.com/argo-agorshechnikov/gRPC-service/proto"
	"google.golang.org/grpc"
)

// Server must be compiled at absent methods
type Server struct {
	// base variant
	pb.UnimplementedExampleServiceServer
}

// Context for timeout and abort
func (s *Server) GetData(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// Reterned response string
	return &pb.Response{Result: "Response on: " + req.Text}, nil
}

func main() {

	// Listening port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Create instance gRPC server
	s := grpc.NewServer()

	// Registration server
	pb.RegisterExampleServiceServer(s, &Server{})

	// Server start
	log.Println("Server start on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
