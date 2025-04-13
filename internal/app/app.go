package app

import (
	"log/slog"
	"os"

	grpcapp "github.com/dinoagera/api-db/internal/app/grpc"
	workdb "github.com/dinoagera/api-db/internal/services/workDB"
	"github.com/dinoagera/api-db/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort string, storagePath string) *App {
	storage, err := postgres.New(storagePath)
	if err != nil {
		log.Debug("storage is not init, on", "path:", storagePath)
		os.Exit(1)
	}
	workService := workdb.New(log, storage, storage, storage, storage)
	grpcApp := grpcapp.New(log, workService, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}
