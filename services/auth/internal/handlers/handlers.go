package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/auth/internal/services"
	"github.com/popeskul/awesome-messanger/services/auth/pb/proto"
)

type Services interface {
	AuthService() services.AuthServiceI
	TokenService() services.TokenServiceI
}

type Validator interface {
	Struct(interface{}) error
}

type Handler struct {
	proto.UnimplementedAuthServiceServer
	services  Services
	validator Validator
}

func NewHandlers(services Services, validator Validator) *Handler {
	return &Handler{
		services:  services,
		validator: validator,
	}
}

func (s *Handler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	input := LoginRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}

	token, err := s.services.AuthService().Login(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Token: token}, nil
}

func (s *Handler) Logout(ctx context.Context, req *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	input := LogoutRequest{
		Token: req.GetToken(),
	}

	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}

	message, err := s.services.AuthService().Logout(ctx, input.Token)
	if err != nil {
		return nil, err
	}

	return &proto.LogoutResponse{Message: message}, nil
}

func (s *Handler) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	input := RegisterRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}

	message, err := s.services.AuthService().Register(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return &proto.RegisterResponse{Message: message}, nil
}

func (s *Handler) Refresh(ctx context.Context, req *proto.RefreshRequest) (*proto.RefreshResponse, error) {
	input := RefreshRequest{
		OldToken: req.GetOldToken(),
	}

	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}

	newToken, err := s.services.AuthService().Refresh(ctx, req.GetOldToken())
	if err != nil {
		return nil, err
	}
	return &proto.RefreshResponse{NewToken: newToken}, nil
}
