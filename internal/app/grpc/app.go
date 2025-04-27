package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	apidb "github.com/dinoagera/api-db/internal/grpc/api-db"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func New(log *slog.Logger, workService apidb.WorkDB, port string) *App {
	gRPCServer := grpc.NewServer()
	apidb.Register(gRPCServer, workService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))
	if err != nil {
		a.log.Debug("start GRPC server to failed on", "port:", a.port)
		a.log.Info("start GRPC server to failed")
		return fmt.Errorf("error:%s", err.Error())
	}
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("error:%s", err.Error())
	}
	a.log.Debug("starting GRPC server on", "address", l.Addr().String())
	a.log.Info("starting GRPC server")
	return nil
}
func (a *App) Stop() {
	a.log.Info("grpc server stopped")
	a.gRPCServer.GracefulStop()
}
