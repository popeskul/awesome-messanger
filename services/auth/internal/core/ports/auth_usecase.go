package ports

import (
	"context"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type RegisterResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type AuthUseCase interface {
	VerifyCredentials(ctx context.Context, username, password string) error
	Logout(ctx context.Context, token string) error
	Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error)
	Refresh(ctx context.Context, token string) (string, error)
	Me(ctx context.Context, token string) (string, error)
}
