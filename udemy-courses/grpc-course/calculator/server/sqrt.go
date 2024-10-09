package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-course/calculator/proto"
	"log"
	"math"
)

// Sqrt Here we will handle errors in the server side
func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with: %v\n", in)

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
