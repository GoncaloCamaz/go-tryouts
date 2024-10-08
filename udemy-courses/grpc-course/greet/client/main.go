package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-course/greet/proto"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	doLongGreet(c)
}
