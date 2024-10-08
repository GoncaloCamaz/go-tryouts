package main

import (
	"fmt"
	pb "grpc-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Failed to receive request: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
