package handlers

import (
	"context"
	"fmt"

	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/grpc/converters"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
	"github.com/popeskul/awesome-messanger/services/profile/proto/api/grpcutils"
)

type profileHandler struct {
	profile.UnimplementedProfileServiceServer

	profileUseCase ports.ProfileUseCase
	validator      ports.Validator
	logger         ports.Logger
}

func NewMessageHandler(
	profileUseCase ports.ProfileUseCase,
	validator ports.Validator,
	logger ports.Logger,
) profile.ProfileServiceServer {
	return &profileHandler{
		profileUseCase: profileUseCase,
		validator:      validator,
		logger:         logger,
	}
}

func (h *profileHandler) CreateProfile(ctx context.Context, request *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	_, err := h.profileUseCase.CreateProfile(ctx, converters.ProtoCreateProfileRequestToPortsRegisterRequest(request))
	if err != nil {
		h.logger.Error("failed to create profile", "error", err)
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}

	return &profile.CreateProfileResponse{
		Success: true,
		Message: "Profile created",
	}, nil
}

func (h *profileHandler) UpdateProfile(ctx context.Context, request *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	_, err := h.profileUseCase.UpdateProfile(ctx, converters.ProtoUpdateProfileRequestToPortsUpdateRequest(request))
	if err != nil {
		h.logger.Error("failed to update profile", "error", err)
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return &profile.UpdateProfileResponse{
		Success: true,
		Message: "Profile updated",
	}, nil
}

func (h *profileHandler) GetProfile(ctx context.Context, request *profile.GetProfileRequest) (*profile.GetProfileResponse, error) {
	if err := h.validator.Validate(request); err != nil {
		h.logger.Error("validation error", "error", err)
		return nil, grpcutils.RPCValidationError(err)
	}

	result, err := h.profileUseCase.GetProfile(ctx, converters.ProtoGetProfileRequestToPortsGetRequest(request))
	if err != nil {
		h.logger.Error("failed to get profile", "error", err)
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return converters.DomainGetProfileResponseToProtoGetProfileResponse(result), nil
}
