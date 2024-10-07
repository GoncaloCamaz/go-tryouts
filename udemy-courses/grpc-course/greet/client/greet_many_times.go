package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	res := &pb.GreetRequest{
		FirstName: "Kelvin",
	}

	stream, err := c.GreetManyTimes(context.Background(), res)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}
