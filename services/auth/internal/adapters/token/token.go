package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
)

var jwtSecret = []byte("your-secret-key")

type TokenService struct{}

func NewTokenService() ports.TokenManager {
	return &TokenService{}
}

func (s *TokenService) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenService) ValidateToken(tokenString string) (string, error) {
	// Implement token validation logic here
	return "", nil
}
