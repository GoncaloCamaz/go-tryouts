package repository

import (
	pb "class-app/internal/api-class/handlers/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Server struct {
	pb.ClassServiceServer
}

func GetClassInfo(classId int64) error {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
		return err
	}

	ctx := context.Background()
	c := pb.NewClassServiceClient(conn)

	req := &pb.ClassRequest{
		ClassId: classId,
	}
	classInfo, err := c.GetClassInfo(ctx, req)

	if err != nil {
		log.Fatalf("Failed to get class info: %v\n", err)
		return err
	}

	log.Printf("Class info: %v\n", classInfo)
	return nil
}
