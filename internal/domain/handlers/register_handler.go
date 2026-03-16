package handlers

import (
	"context"
	"errors"
	"log/slog"
	v1 "sc_auth/generated/skycontrol/proto/auth/v1"
	db_gen "sc_auth/internal/db/gen"
	"sc_auth/internal/domain/models"
)

func RegisterHandler(
	ctx context.Context,
	req *v1.RegisterRequest,
	logger *slog.Logger,
	db db_gen.Queries,
) (*v1.RegisterResponse, error) {
	logger.Debug("register request received", "request", req)

	isExists, err := db.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if isExists {
		return &v1.RegisterResponse{
			ErrMessage: "A user with this email address has already been created",
		}, nil
	}
	logger.Debug(" ", "request", req)
	if isExists, err := db.FindUserByUsername(ctx, req.Username); err != nil {
		return nil, err
	} else if isExists {
		return &v1.RegisterResponse{
			ErrMessage: "A user with this username has already been created",
		}, nil
	}

	user, err := models.NewUser(req.Email, req.Password, req.Username)
	if err != nil {
		return nil, errors.New("failed to create user model")
	}
	if err := db.CreateUser(ctx, db_gen.CreateUserParams{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PassHash,
		Username:     user.Username,
	}); err != nil {
		return nil, err
	}

	return &v1.RegisterResponse{
		UserId:     user.ID,
		Token:      "Заглушка пока что",
		ErrMessage: "",
	}, nil
}
