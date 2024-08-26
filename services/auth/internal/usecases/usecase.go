package usecases

import (
	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
)

type UseCase struct {
	auth  ports.AuthUseCase
	token ports.TokenUseCase
}

func NewUseCase(
	authUseCase ports.AuthUseCase,
	tokenUseCase ports.TokenUseCase,
) *UseCase {
	return &UseCase{
		auth:  authUseCase,
		token: tokenUseCase,
	}
}

func (u *UseCase) AuthUseCase() ports.AuthUseCase {
	return u.auth
}

func (u *UseCase) TokenUseCase() ports.TokenUseCase {
	return u.token
}
