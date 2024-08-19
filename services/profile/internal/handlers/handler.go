package handlers

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/popeskul/awesome-messanger/services/profile/internal/model"
	"github.com/popeskul/awesome-messanger/services/profile/internal/services"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

type Services interface {
	ProfileService() services.ProfileServiceI
}

type Handler struct {
	profile.UnimplementedProfileServiceServer
	health.UnimplementedHealthServiceServer
	services  Services
	validator *protovalidate.Validator
}

func NewHandler(services Services) (*Handler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&profile.CreateProfileRequest{},
			&profile.UpdateProfileRequest{},
			&profile.GetProfileRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	return &Handler{
		services:  services,
		validator: validator,
	}, err
}

func (h *Handler) CreateProfile(ctx context.Context, req *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, err
	}

	input := &model.CreateProfileRequest{
		UserId:    req.GetUserId(),
		Nickname:  req.GetNickname(),
		Bio:       req.GetBio(),
		AvatarUrl: req.GetAvatarUrl(),
	}

	if _, err := h.services.ProfileService().CreateProfile(ctx, input); err != nil {
		return nil, err
	}

	return &profile.CreateProfileResponse{
		Success: true,
		Message: "Profile created",
	}, nil
}

func (h *Handler) UpdateProfile(ctx context.Context, req *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, err
	}

	input := &model.UpdateProfileRequest{
		UserId:    req.GetUserId(),
		Nickname:  req.GetNickname(),
		Bio:       req.GetBio(),
		AvatarUrl: req.GetAvatarUrl(),
	}

	if _, err := h.services.ProfileService().UpdateProfile(ctx, input); err != nil {
		return nil, err
	}

	return &profile.UpdateProfileResponse{
		Success: true,
		Message: "Profile updated",
	}, nil
}

func (h *Handler) GetProfile(ctx context.Context, req *profile.GetProfileRequest) (*profile.GetProfileResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, err
	}

	input := &model.GetProfileRequest{
		UserId: req.GetUserId(),
	}

	res, err := h.services.ProfileService().GetProfile(ctx, input)
	if err != nil {
		return nil, err
	}

	return &profile.GetProfileResponse{
		UserId:    res.UserId,
		Nickname:  res.Nickname,
		Bio:       res.Bio,
		AvatarUrl: res.AvatarUrl,
	}, nil
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
