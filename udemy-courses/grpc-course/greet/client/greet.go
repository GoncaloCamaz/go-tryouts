package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Kelvin"})

	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}

	log.Printf("Greeting response: %v\n", res.Result)
}
