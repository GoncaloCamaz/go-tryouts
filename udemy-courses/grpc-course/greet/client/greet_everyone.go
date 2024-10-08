package main

import (
	"context"
	pb "grpc-course/greet/proto"
	"io"
	"log"
	"sync"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Failed to greet everyone: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Kelvin"},
		{FirstName: "John"},
		{FirstName: "Jane"},
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Sends a bunch of messages to the server (go routine)
	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
		wg.Done()
	}()

	// Receives a bunch of messages from the server (go routine)
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Failed to receive response: %v\n", err)
				break
			}

			log.Printf("Received response: %v\n", res.GetResult())
		}

		wg.Done()
	}()
	wg.Wait()
}
