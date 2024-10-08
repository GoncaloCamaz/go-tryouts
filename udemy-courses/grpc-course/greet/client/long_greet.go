package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kelvin"},
		{FirstName: "Marie"},
		{FirstName: "John"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed to long greet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req.FirstName)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Failed to receive response: %v\n", err)
	}

	log.Printf("LongGreet response: %v\n", res.Result)
}
