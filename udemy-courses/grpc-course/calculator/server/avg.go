package main

import (
	pb "grpc-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average was invoked")

	res := 0
	numberOfRequests := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(res) / float64(numberOfRequests),
			})
		}

		if err != nil {
			log.Fatalf("Failed to receive request: %v\n", err)
		}

		numberOfRequests++
		res += int(req.Numbers)
	}
}
