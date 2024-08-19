package services

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/profile/internal/model"
)

type ProfileService struct {
	logger Logger
}

func NewProfileService(logger Logger) *ProfileService {
	return &ProfileService{
		logger: logger,
	}
}

func (s *ProfileService) UpdateProfile(ctx context.Context, req *model.UpdateProfileRequest) (*model.Profile, error) {
	s.logger.Printf("UpdateProfile called")

	return &model.Profile{
		UserId:    req.UserId,
		Nickname:  req.Nickname,
		Bio:       req.Bio,
		AvatarUrl: req.AvatarUrl,
	}, nil
}

func (s *ProfileService) GetProfile(ctx context.Context, req *model.GetProfileRequest) (*model.Profile, error) {
	s.logger.Printf("GetProfile called")

	return &model.Profile{
		UserId:    req.UserId,
		Nickname:  "nickname",
		Bio:       "bio",
		AvatarUrl: "avatar_url",
	}, nil
}

func (s *ProfileService) CreateProfile(ctx context.Context, req *model.CreateProfileRequest) (*model.Profile, error) {
	s.logger.Printf("CreateProfile called")

	return &model.Profile{
		UserId:    req.UserId,
		Nickname:  req.Nickname,
		Bio:       req.Bio,
		AvatarUrl: req.AvatarUrl,
	}, nil
}
