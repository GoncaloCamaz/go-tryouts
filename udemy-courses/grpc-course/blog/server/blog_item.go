package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	pb "grpc-course/blog/proto"
)

type BlogItem struct {
	// this is the id of the document in the database, we add omitempty so it doesn't show up in the response
	// omitempty is a tag that tells the bson package to not include the field in the document if it's empty
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		// this .Hex transforms the id into a string
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
