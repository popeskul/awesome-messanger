package service

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/search/internal/graph/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) FindUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	return &model.User{
		ID:          "1",
		Nickname:    nickname,
		Description: "Static description for search result",
	}, nil
}
