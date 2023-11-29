package main

import (
	"context"
	"fmt"
	desc "github.com/Task_tracker_back/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

// GET ...
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetResponse{
		Info: &desc.UserInfo{
			Id:      req.GetId(),
			Name:    "Anton",
			IsHuman: false,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %d", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %d", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %d", err)
	}
}
