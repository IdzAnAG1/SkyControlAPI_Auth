package app

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	authv1 "sc_auth/generated/skycontrol/proto/auth/v1"
	"sc_auth/internal/config"
	"sc_auth/internal/interruptor"
	logger "sc_auth/internal/logger"
	"sc_auth/internal/server"

	"google.golang.org/grpc"
)

/*
	Entry point for Microservice
*/

type App struct {
	tcpPort string
	logger  *slog.Logger
}

func NewApp() (*App, error) {
	cfg, err := config.LoadAndGetConfig("~/sc_auth.env")
	if err != nil {
		return nil, err
	}

	return &App{
		tcpPort: cfg.AuthServer.Port,
		logger:  logger.New(),
	}, nil
}

func (app *App) Run() error {
	app.logger.Info("Starting Auth Server", "port", app.tcpPort)
	// todo
	tcp, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%s", app.tcpPort),
	)
	if err != nil {
		return err
	}
	defer func() {
		err = tcp.Close()

		if err == nil || errors.Is(err, net.ErrClosed) {
			app.logger.Info("tcp listener is closed")
			return
		}

		app.logger.Error("closing tcp listener is failed", "error", err)
	}()

	// todo
	srv := grpc.NewServer()
	app.logger.Info("grpc server is created")
	// todo
	iter := interruptor.NewInterruptor(srv, app.logger)
	iter.Run()

	// todo
	authServ := server.GrpcAuthServer{}
	authv1.RegisterAuthServer(srv, &authServ)
	app.logger.Info("auth server is registered")
	err = srv.Serve(tcp)
	if err != nil {
		return err
	}
	return nil
}
