package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dinoagera/api-db/internal/app"
	"github.com/dinoagera/api-db/internal/config"
	"github.com/dinoagera/api-db/internal/logger"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.InitConfig(logger)
	application := app.New(logger, cfg.GRPC.Port, cfg.StoragePath)
	go application.GRPCServer.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GRPCServer.Stop()
	logger.Info("application is stopped")
}
