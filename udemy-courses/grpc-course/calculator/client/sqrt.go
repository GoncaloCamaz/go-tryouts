package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-course/calculator/proto"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		// lets try to transform error in grpc status
		e, ok := status.FromError(err)

		// if ok, we have a grpc status error
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Printf("Error code from server: %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number\n")
				return
			}

		} else {
			log.Fatalf("Failed to sqrt: %v\n", err)
		}
	}

	log.Printf("Sqrt response: %v\n", res.Result)
}
