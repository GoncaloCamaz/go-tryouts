package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-course/greet/proto"
	"log"
	"time"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change this to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed to load certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// use line bellow if tls false
	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// use line bellow if tls true
	conn, err := grpc.Dial(addr, opts...)
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	// Unary RPC, client sends one request and server returns one response
	//doGreet(c)
	// Server Streaming, client sends one request
	//doGreetManyTimes(c)
	// Client Streaming, server returns one response
	//doLongGreet(c)
	// Bi-Directional Streaming
	//doGreetEveryone(c)

	// do greet with timeout
	doGreetWithDeadline(c, 5*time.Second)
}
