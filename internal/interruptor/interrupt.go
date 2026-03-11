package interruptor

import (
	"log/slog"
	"os"
	"os/signal"
	"sc_auth/internal/db"
	"syscall"

	"google.golang.org/grpc"
)

type Interruptor struct {
	gRPCInterruptor *grpc.Server
	signal          chan os.Signal
	logger          *slog.Logger
	database        db.DB
}

func NewInterruptor(srv *grpc.Server, logger *slog.Logger, db db.DB) *Interruptor {
	return &Interruptor{
		srv,
		make(chan os.Signal, 1),
		logger,
		db,
	}
}

func (i *Interruptor) Run() {
	i.startCatchingSignal()
	go func() {
		i.shutdown()
	}()
}

func (i *Interruptor) startCatchingSignal() {
	i.logger.Info("Starting signal catching")
	signal.Notify(i.signal, syscall.SIGTERM, syscall.SIGINT)
}

func (i *Interruptor) shutdown() {
	<-i.signal
	i.logger.Info("Server is shutting down gracefully...")
	i.gRPCInterruptor.GracefulStop()

	err := i.database.Close()
	if err != nil {
		i.logger.Error("Error closing database", "error", err)
	}
}
