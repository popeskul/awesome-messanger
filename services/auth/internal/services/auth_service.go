package services

import (
	"context"
	"github.com/popeskul/awesome-messanger/services/auth/internal/token"
)

type AuthService struct {
	logger Logger
}

func NewAuthService(logger Logger) *AuthService {
	return &AuthService{
		logger: logger,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	s.logger.Printf("Login request received for user %s", username)

	generateToken, err := token.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return generateToken, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) (string, error) {
	s.logger.Printf("Logout request received for token %s", token)

	return "Logged out", nil
}

func (s *AuthService) Register(ctx context.Context, username, password string) (string, error) {
	s.logger.Printf("Register request received for user %s", username)

	return "Registered", nil
}

func (s *AuthService) Refresh(ctx context.Context, oldToken string) (string, error) {
	s.logger.Printf("Refresh request received for old token %s", oldToken)

	return "", nil
}
