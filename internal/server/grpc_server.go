package server

import (
	"context"
	v1 "sc_auth/generated/skycontrol/proto/auth/v1"
)

type GrpcAuthServer struct {
	v1.UnimplementedAuthServer
}

func (gs *GrpcAuthServer) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return nil, nil
}

func (gs *GrpcAuthServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return nil, nil
}
