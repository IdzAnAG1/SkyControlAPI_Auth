package interruptor

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type Interruptor struct {
	gRPCInterruptor *grpc.Server
	signal          chan os.Signal
}

func NewInterruptor(srv *grpc.Server) *Interruptor {
	return &Interruptor{
		srv,
		make(chan os.Signal, 1),
	}
}

func (i *Interruptor) Run() (err error) {
	i.startCatchingSignal()
	go func() {
		i.shutdown()
	}()
	return nil
}

func (i *Interruptor) startCatchingSignal() {
	signal.Notify(i.signal, syscall.SIGTERM, syscall.SIGINT)
}

func (i *Interruptor) shutdown() {
	<-i.signal
	fmt.Println("\nServer is shutting down gracefully...")
	i.gRPCInterruptor.GracefulStop()
}
