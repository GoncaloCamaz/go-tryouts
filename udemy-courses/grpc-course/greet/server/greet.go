package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
