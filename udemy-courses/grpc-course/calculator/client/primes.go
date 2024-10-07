package main

import (
	"context"
	pb "grpc-course/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimeRequest{Number: 123}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling primes!")
	}

	for {
		res, err := stream.Recv()

		if err != io.EOF {
			return
		}

		if err != nil {
			log.Fatalf("error while reading stream %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
