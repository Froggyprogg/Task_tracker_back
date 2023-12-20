package app

import (
	"fmt"
	"github.com/Task_tracker_back/pkg/config"
	"github.com/Task_tracker_back/pkg/db"
	"github.com/Task_tracker_back/pkg/grpc/board"
	"github.com/Task_tracker_back/pkg/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	database   *gorm.DB
}

func Run(cfg *config.Config) *App {
	database := db.NewDatabaseConnection(cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %d", err)
		return nil
	}

	gRPCServer := grpc.NewServer()
	reflection.Register(gRPCServer)
	user.Register(gRPCServer, database)
	board.Register(gRPCServer, database)

	log.Printf("server listening at %d", lis.Addr())

	if err = gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %d", err)
		return nil
	}

	return &App{gRPCServer: gRPCServer, database: database}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
