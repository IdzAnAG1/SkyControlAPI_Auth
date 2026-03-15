package server

import (
	"context"
	"log/slog"
	v1 "sc_auth/generated/skycontrol/proto/auth/v1"
	"sc_auth/internal/db"
)

type GrpcAuthServer struct {
	v1.UnimplementedAuthServer
	Logger *slog.Logger
	DB     *db.DB
}

func (gs *GrpcAuthServer) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	gs.Logger.Debug("register request received", "request", req)
	return nil, nil
}

func (gs *GrpcAuthServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	gs.Logger.Debug("login request received", "request", req)
	return nil, nil
}
