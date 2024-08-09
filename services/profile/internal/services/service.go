package services

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/profile/internal/model"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type ProfileServiceI interface {
	UpdateProfile(ctx context.Context, req *model.UpdateProfileRequest) (*model.Profile, error)
	GetProfile(ctx context.Context, req *model.GetProfileRequest) (*model.Profile, error)
}

type IServices interface {
	ProfileService() ProfileServiceI
}

type Services struct {
	profileService ProfileServiceI
}

func NewService(profileService ProfileServiceI) IServices {
	return &Services{
		profileService: profileService,
	}
}

func (s *Services) ProfileService() ProfileServiceI {
	return s.profileService
}
