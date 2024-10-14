package grpc

import (
	ClassDataModel "class-app/internal/api-class/datamodel"
	pb "class-app/internal/api-class/handlers/grpc/proto"
	"context"
	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"time"
)

type Server struct {
	pb.ClassServiceServer
	DB *pg.DB
}

func StartServer(addr string, db *pg.DB) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterClassServiceServer(s, &Server{
		DB: db,
	})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func (s *Server) GetClassInfo(ctx context.Context, in *pb.ClassRequest) (*pb.ClassResponse, error) {
	log.Printf("GetClassInfo function was invoked with: %v\n", in)

	var class ClassDataModel.Class
	err := s.DB.Model(&class).Where("id = ?", in.ClassId).Select()

	if err != nil {
		log.Printf("Database error: %v\n", err)
		return nil, err
	}

	return &pb.ClassResponse{
		Id:      int32(class.ID),
		Number:  int32(class.Number),
		Year:    class.Year,
		Created: class.Created.Format(time.RFC3339),
		Updated: class.Updated.Format(time.RFC3339),
	}, nil
}

func (s *Server) GetClassList(_ context.Context, _ *emptypb.Empty) (*pb.ClassListResponse, error) {
	log.Printf("GetClassList function was invoked\n")

	var classes []ClassDataModel.Class
	err := s.DB.Model(&classes).Select()

	if err != nil {
		log.Printf("Database error: %v\n", err)
		return nil, err
	}

	var classList []*pb.ClassResponse
	for _, class := range classes {
		classList = append(classList, &pb.ClassResponse{
			Id:      int32(class.ID),
			Number:  int32(class.Number),
			Year:    class.Year,
			Created: class.Created.Format(time.RFC3339),
			Updated: class.Updated.Format(time.RFC3339),
		})
	}

	return &pb.ClassListResponse{Classes: classList}, nil
}
