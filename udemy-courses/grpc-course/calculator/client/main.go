package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-course/calculator/proto"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}

	c := pb.NewCalculatorServiceClient(conn)

	//calculateAverage(c)
	doSqrt(c, 10)
}
