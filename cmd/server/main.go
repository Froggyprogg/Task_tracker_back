package main

import (
	"github.com/Task_tracker_back/pkg/app"
	"github.com/Task_tracker_back/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	application := app.Run(cfg)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	application.Stop()
	log.Print("Gracefully stopped")
}
