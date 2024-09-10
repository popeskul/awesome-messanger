package usecases

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
)

type AuthUseCase struct {
	Logger ports.Logger
}

func NewAuthUseCase(logger ports.Logger) ports.AuthUseCase {
	return &AuthUseCase{
		Logger: logger,
	}
}

func (uc *AuthUseCase) VerifyCredentials(ctx context.Context, username, password string) error {
	uc.Logger.Info("VerifyCredentials request received for username %s", username)

	return nil
}

func (uc *AuthUseCase) Logout(ctx context.Context, token string) error {
	uc.Logger.Info("Logout request received for token %s", token)

	return nil
}

func (uc *AuthUseCase) Register(ctx context.Context, user ports.RegisterRequest) (ports.RegisterResponse, error) {
	uc.Logger.Info("Register request received for user %v", user)

	return ports.RegisterResponse{
		User: ports.UserResponse{
			ID:       "id",
			Username: user.Username,
			Email:    user.Email,
		},
		Token: "uuid token",
	}, nil
}

func (uc *AuthUseCase) Refresh(ctx context.Context, oldToken string) (string, error) {
	uc.Logger.Info("Refresh request received for old token %s", oldToken)

	return "token", nil
}

func (uc *AuthUseCase) Me(ctx context.Context, token string) (string, error) {
	uc.Logger.Info("Me request received for token %s", token)

	return "Me", nil
}
