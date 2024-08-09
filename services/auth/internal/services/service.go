package services

import "context"

type Logger interface {
	Printf(format string, v ...interface{})
}

type AuthServiceI interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, password string) (string, error)
	Logout(ctx context.Context, token string) (string, error)
	Refresh(ctx context.Context, oldToken string) (string, error)
}

type TokenServiceI interface {
	GenerateToken(email string) (string, error)
}

type IServices interface {
	AuthService() AuthServiceI
	TokenService() TokenServiceI
}

type Services struct {
	tokenService TokenServiceI
	authService  AuthServiceI
}

func NewService(tokenService TokenServiceI, authService AuthServiceI) IServices {
	return &Services{
		tokenService: tokenService,
		authService:  authService,
	}
}

func (s *Services) AuthService() AuthServiceI {
	return s.authService
}

func (s *Services) TokenService() TokenServiceI {
	return s.tokenService
}
