package server

import (
	"context"
	"log/slog"
	v1 "sc_auth/generated/skycontrol/proto/auth/v1"
	"sc_auth/internal/db"
	"sc_auth/internal/domain/handlers"
)

type GrpcAuthServer struct {
	v1.UnimplementedAuthServer
	Logger *slog.Logger
	DB     *db.DB
}

func (gs *GrpcAuthServer) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return handlers.RegisterHandler(ctx, req, gs.Logger, *gs.DB.Queries)
}

func (gs *GrpcAuthServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return handlers.LoginHandler(ctx, req, gs.Logger, *gs.DB.Queries)
}
