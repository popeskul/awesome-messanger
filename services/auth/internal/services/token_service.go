package services

import "github.com/popeskul/awesome-messanger/services/auth/internal/token"

type TokenService struct {
	logger Logger
}

func NewTokenService(logger Logger) *TokenService {
	return &TokenService{
		logger: logger,
	}
}

func (s *TokenService) GenerateToken(email string) (string, error) {
	generateToken, err := token.GenerateToken(email)
	if err != nil {
		s.logger.Printf("Error generating token: %v", err)
		return "", err
	}

	return generateToken, nil
}
