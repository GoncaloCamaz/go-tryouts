package main

import (
	"context"
	pb "grpc-course/calculator/proto"
	"log"
	"time"
)

func calculateAverage(c pb.CalculatorServiceClient) {
	log.Println("calculateAverage was invoked")

	reqs := []*pb.AvgRequest{
		{Numbers: 1},
		{Numbers: 2},
		{Numbers: 3},
		{Numbers: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Failed to calculate average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req.Numbers)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Failed to receive response: %v\n", err)
	}

	log.Printf("Average response: %v\n", res.Result)
}
