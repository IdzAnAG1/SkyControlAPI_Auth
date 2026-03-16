package handlers

import (
	"context"
	"errors"
	"log/slog"
	v1 "sc_auth/generated/skycontrol/proto/auth/v1"
	db_gen "sc_auth/internal/db/gen"
	"sc_auth/internal/domain/utils"

	"github.com/jackc/pgx/v5"
)

func LoginHandler(
	ctx context.Context,
	req *v1.LoginRequest,
	logger *slog.Logger,
	db db_gen.Queries,
) (*v1.LoginResponse, error) {
	logger.Debug("login request received", "request", req)

	user, err := db.GetUserByEmail(ctx, req.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		return &v1.LoginResponse{ErrMessage: "User not found"}, nil
	} else if err != nil {
		return nil, err
	}

	err = utils.CheckPasswordHash(user.PasswordHash, req.Password)
	if err != nil {
		return &v1.LoginResponse{ErrMessage: "Invalid password"}, nil
	}

	return &v1.LoginResponse{UserId: user.ID, Token: "token"}, nil
}
