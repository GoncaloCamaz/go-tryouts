package main

import (
	"context"
	pb "grpc-course/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Create blog invoked")
	blog := &pb.Blog{
		AuthorId: "Camaz do Aço",
		Title:    "My go tryouts",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog created: %v\n", res)
	return res.Id
}
