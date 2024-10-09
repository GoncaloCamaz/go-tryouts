package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-course/greet/proto"
	"log"
	"net"
)

var addr string = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	tls := true // change this to false if needed
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		key := "ssl/server.pem"
		creds, err := credentials.NewClientTLSFromFile(certFile, key)

		if err != nil {
			log.Fatalf("Failed to load certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
