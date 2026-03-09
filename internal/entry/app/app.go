package app

import (
	"fmt"
	"log"
	"net"
	authv1 "sc_auth/generated/skycontrol/proto/auth/v1"
	"sc_auth/internal/interruptor"
	"sc_auth/internal/server"
	"strconv"

	"google.golang.org/grpc"
)

/*
	Entry point for Microservice
*/

type App struct {
	tcpPort int
	log     *log.Logger
}

func NewApp(tcp int) *App {
	return &App{
		tcpPort: tcp,
	}
}

func (app *App) Run() error {
	// todo
	tcp, err := net.Listen("tcp", ":"+strconv.Itoa(app.tcpPort))
	if err != nil {
		return err
	}
	defer func() {
		err = tcp.Close()
		if err != nil {
			fmt.Println("tcp Server is down")
		}
	}()

	// todo
	srv := grpc.NewServer()

	// todo
	iter := interruptor.NewInterruptor(srv)
	err = iter.Run()
	if err != nil {
		return err
	}

	// todo
	authServ := server.GrpcAuthServer{}
	authv1.RegisterAuthServer(srv, &authServ)
	err = srv.Serve(tcp)
	if err != nil {
		return err
	}
	return nil
}
