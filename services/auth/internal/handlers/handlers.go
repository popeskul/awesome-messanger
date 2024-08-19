package handlers

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/popeskul/awesome-messanger/services/auth/internal/services"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/grpcutils"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
)

type Services interface {
	AuthService() services.AuthServiceI
	TokenService() services.TokenServiceI
}

type Handler struct {
	auth.UnimplementedAuthServiceServer
	health.UnimplementedHealthServiceServer
	services  Services
	validator *protovalidate.Validator
}

func NewHandlers(services Services) (*Handler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&auth.RegisterRequest{},
			&auth.LoginRequest{},
			&auth.RefreshRequest{},
			&auth.LogoutRequest{},
			&health.HealthCheckRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	return &Handler{
		services:  services,
		validator: validator,
	}, nil
}

func (h *Handler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	token, err := h.services.AuthService().Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{Token: token}, nil
}

func (h *Handler) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	message, err := h.services.AuthService().Logout(ctx, req.GetToken())
	if err != nil {
		return nil, err
	}

	return &auth.LogoutResponse{Message: message}, nil
}

func (h *Handler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	message, err := h.services.AuthService().Register(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{Message: message}, nil
}

func (h *Handler) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	newToken, err := h.services.AuthService().Refresh(ctx, req.GetOldToken())
	if err != nil {
		return nil, err
	}
	return &auth.RefreshResponse{NewToken: newToken}, nil
}

func (h *Handler) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "healthy"}, nil
}

func (h *Handler) Liveness(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "alive"}, nil
}

func (h *Handler) Readiness(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "ready"}, nil
}

func (h *Handler) Healthz(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "healthy"}, nil
}
