package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/profile/internal/model"
	"github.com/popeskul/awesome-messanger/services/profile/internal/services"
	"github.com/popeskul/awesome-messanger/services/profile/pb/proto"
)

type Services interface {
	ProfileService() services.ProfileServiceI
}

type Validator interface {
	Struct(interface{}) error
}

type Handler struct {
	proto.UnimplementedProfileServiceServer
	services  Services
	validator Validator
}

func NewHandler(services Services, validator Validator) *Handler {
	return &Handler{
		services:  services,
		validator: validator,
	}
}

func (h *Handler) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	input := &model.UpdateProfileRequest{
		UserId:    req.GetUserId(),
		Nickname:  req.GetNickname(),
		Bio:       req.GetBio(),
		AvatarUrl: req.GetAvatarUrl(),
	}

	if err := h.validator.Struct(input); err != nil {
		return nil, err
	}

	if _, err := h.services.ProfileService().UpdateProfile(ctx, input); err != nil {
		return nil, err
	}

	return &proto.UpdateProfileResponse{
		Success: true,
		Message: "Profile updated",
	}, nil
}

func (h *Handler) GetProfile(ctx context.Context, req *proto.GetProfileRequest) (*proto.GetProfileResponse, error) {
	input := &model.GetProfileRequest{
		UserId: req.GetUserId(),
	}

	if err := h.validator.Struct(input); err != nil {
		return nil, err
	}

	res, err := h.services.ProfileService().GetProfile(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.GetProfileResponse{
		UserId:    res.UserId,
		Nickname:  res.Nickname,
		Bio:       res.Bio,
		AvatarUrl: res.AvatarUrl,
	}, nil
}
