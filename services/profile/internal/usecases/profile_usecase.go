package usecases

import (
	"context"
	"errors"

	"github.com/popeskul/awesome-messanger/services/profile/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
)

type profileUseCase struct {
	logger ports.Logger
}

func NewProfileUseCase(logger ports.Logger) ports.ProfileUseCase {
	return &profileUseCase{
		logger: logger,
	}
}

func (uc profileUseCase) CreateProfile(ctx context.Context, profile *ports.CreateProfileRequest) (*domain.Profile, error) {
	uc.logger.Info("CreateProfile called")

	if profile.UserId == "" {
		return nil, errors.New("user_id is required")
	} else if profile.Nickname == "" {
		return nil, errors.New("nickname is required")
	}

	return &domain.Profile{
		UserId:    profile.UserId,
		Nickname:  profile.Nickname,
		Bio:       profile.Bio,
		AvatarUrl: profile.AvatarUrl,
	}, nil
}

func (uc profileUseCase) UpdateProfile(ctx context.Context, profile *ports.UpdateProfileRequest) (*domain.Profile, error) {
	uc.logger.Info("UpdateProfile called")

	if profile.UserId == "" {
		return nil, errors.New("invalid user id")
	}
	if profile.Nickname == "" && profile.Bio == "" && profile.AvatarUrl == "" {
		return nil, errors.New("no fields to update")
	}

	return &domain.Profile{
		UserId:    profile.UserId,
		Nickname:  profile.Nickname,
		Bio:       profile.Bio,
		AvatarUrl: profile.AvatarUrl,
	}, nil
}

func (uc profileUseCase) GetProfile(ctx context.Context, profile *ports.GetProfileRequest) (*domain.Profile, error) {
	uc.logger.Info("GetProfile called")

	if profile.UserId == "" {
		return nil, errors.New("invalid user id")
	}

	if profile.UserId == "nonexistent" {
		return nil, errors.New("profile not found")
	}

	return &domain.Profile{
		UserId:    profile.UserId,
		Nickname:  "nickname",
		Bio:       "bio",
		AvatarUrl: "avatar_url",
	}, nil
}
