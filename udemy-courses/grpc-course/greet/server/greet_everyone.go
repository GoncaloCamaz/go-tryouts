package main

import (
	pb "grpc-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed to receive client request: %v\n", err)
			return err
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{Result: res})

		if err != nil {
			log.Fatalf("Failed to send response to client: %v\n", err)
		}
	}
}
