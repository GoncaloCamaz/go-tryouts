package grpc

import (
	pb "class-app/internal/api-class/handlers/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	pb.ClassServiceServer
}

func StartServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterClassServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func (s *Server) GetClassInfo(ctx context.Context, in *pb.ClassRequest) (*pb.ClassResponse, error) {
	log.Printf("GetClassInfo function was invoked with: %v\n", in)
	//todo use repository to get class info
	return &pb.ClassResponse{
		Number:  1,
		Year:    "2023",
		Created: time.Now().Format("YYYY-MM-DD"),
		Updated: time.Now().Format("YYYY-MM-DD"),
	}, nil
}
