package usecases

import (
	"context"

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

	return &domain.Profile{
		UserId:    profile.UserId,
		Nickname:  profile.Nickname,
		Bio:       profile.Bio,
		AvatarUrl: profile.AvatarUrl,
	}, nil
}

func (uc profileUseCase) UpdateProfile(ctx context.Context, profile *ports.UpdateProfileRequest) (*domain.Profile, error) {
	uc.logger.Info("UpdateProfile called")

	return &domain.Profile{
		UserId:    profile.UserId,
		Nickname:  profile.Nickname,
		Bio:       profile.Bio,
		AvatarUrl: profile.AvatarUrl,
	}, nil
}

func (uc profileUseCase) GetProfile(ctx context.Context, profile *ports.GetProfileRequest) (*domain.Profile, error) {
	uc.logger.Info("GetProfile called")

	return &domain.Profile{
		UserId:    profile.UserId,
		Nickname:  "nickname",
		Bio:       "bio",
		AvatarUrl: "avatar_url",
	}, nil
}
