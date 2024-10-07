package main

import (
	"context"
	pb "grpc-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 1, SecondNumber: 1})

	if err != nil {
		log.Fatalf("Failed to call Sum: %v\n", err)
	}

	log.Printf("Sum result: %v\n", res.Result)
}
