package ports

import (
	"context"
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/domain"
)

type ProfileUseCase interface {
	CreateProfile(ctx context.Context, profile *CreateProfileRequest) (*domain.Profile, error)
	UpdateProfile(ctx context.Context, profile *UpdateProfileRequest) (*domain.Profile, error)
	GetProfile(ctx context.Context, profile *GetProfileRequest) (*domain.Profile, error)
}

type UserCase interface {
	ProfileUseCase() ProfileUseCase
}
