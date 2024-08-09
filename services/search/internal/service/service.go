package service

import (
	"context"
	"github.com/popeskul/awesome-messanger/services/search/internal/graph/model"
)

type IUserService interface {
	FindUserByNickname(ctx context.Context, nickname string) (*model.User, error)
}

type IService interface {
	UserService() IUserService
}

type Service struct {
	userService IUserService
}

func NewService(userService IUserService) IService {
	return &Service{
		userService: userService,
	}
}

func (s *Service) UserService() IUserService {
	return s.userService
}
