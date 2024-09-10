package usecases

import (
	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
)

type TokenUseCase struct {
	Logger ports.Logger
	Token  ports.Token
}

func NewTokenUseCase(Logger ports.Logger, Token ports.Token) ports.TokenUseCase {
	return &TokenUseCase{
		Logger: Logger,
		Token:  Token,
	}
}

func (s *TokenUseCase) GenerateToken(email string) (string, error) {
	generateToken, err := s.Token.GenerateToken(email)
	if err != nil {
		s.Logger.Info("Error generating token: %v", err)
		return "", err
	}

	return generateToken, nil
}

func (s *TokenUseCase) ValidateToken(oldToken string) (string, error) {
	validateToken, err := s.Token.ValidateToken(oldToken)
	if err != nil {
		s.Logger.Info("Error validating token: %v", err)
		return "", err
	}

	return validateToken, nil
}
