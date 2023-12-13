package board

import (
	board "github.com/Task_tracker_back/pkg/board_v1"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type server struct {
	board.UnimplementedBoardV1Server
}

var (
	database *gorm.DB
)

func Register(gRPCServer *grpc.Server, db *gorm.DB) {
	board.RegisterBoardV1Server(gRPCServer, &server{})
	database = db
}
