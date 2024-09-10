package handlers

import (
	"context"
	"fmt"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/converters"

	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/grpcutils"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer

	authUseCase  ports.AuthUseCase
	validator    ports.Validator
	logger       ports.Logger
	tokenUseCase ports.TokenUseCase
}

func NewAuthHandler(
	authUseCase ports.AuthUseCase,
	tokenUseCase ports.TokenUseCase,
	validator ports.Validator,
	logger ports.Logger,
) auth.AuthServiceServer {
	return &AuthHandler{
		authUseCase:  authUseCase,
		tokenUseCase: tokenUseCase,
		validator:    validator,
		logger:       logger,
	}
}

func (h *AuthHandler) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	if err := h.authUseCase.VerifyCredentials(ctx, request.GetUsername(), request.GetPassword()); err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	token, err := h.tokenUseCase.GenerateToken(request.GetUsername())
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &auth.LoginResponse{Token: token}, nil
}

func (h *AuthHandler) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	if err := h.authUseCase.Logout(ctx, request.GetToken()); err != nil {
		return nil, fmt.Errorf("failed to logout: %w", err)
	}

	return &auth.LogoutResponse{}, nil
}

func (h *AuthHandler) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	portsRequest := converters.GRPCRegisterRequestToPortsRegisterRequest(request)
	result, err := h.authUseCase.Register(ctx, portsRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return converters.PostRegisterResponseToGRPCRegisterResponse(result), nil
}

func (h *AuthHandler) Refresh(ctx context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	token, err := h.authUseCase.Refresh(ctx, request.GetOldToken())
	if err != nil {
		return nil, fmt.Errorf("failed to refresh: %w", err)
	}

	return &auth.RefreshResponse{NewToken: token}, nil
}
