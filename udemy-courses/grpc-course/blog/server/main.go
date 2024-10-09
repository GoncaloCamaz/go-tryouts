package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	pb "grpc-course/blog/proto"
	"log"
	"net"
)

var collection *mongo.Collection
var addr = "localhost:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// establish connection to database client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Failed to create new mongoDB client: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	// make server listen to localhost on port 50051
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Server listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
