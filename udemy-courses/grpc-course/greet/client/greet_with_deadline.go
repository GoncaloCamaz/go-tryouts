package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-course/greet/proto"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{FirstName: "Kelvin"})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout was hit! Deadline was exceeded\n")
				return
			} else {
				log.Fatalf("Failed to greet: %v\n", err)
			}
		} else {
			log.Fatalf("Failed to greet: %v\n", err)
		}
	}

	log.Printf("Greeting response: %v\n", res.Result)
}
