package usecases

import (
	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
)

type UseCases struct {
	auth  ports.AuthUseCase
	token ports.TokenUseCase
}

func NewUseCases(
	authUseCase ports.AuthUseCase,
	tokenUseCase ports.TokenUseCase,
) *UseCases {
	return &UseCases{
		auth:  authUseCase,
		token: tokenUseCase,
	}
}

func (u *UseCases) AuthUseCase() ports.AuthUseCase {
	return u.auth
}

func (u *UseCases) TokenUseCase() ports.TokenUseCase {
	return u.token
}
