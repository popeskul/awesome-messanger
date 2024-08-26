package usecase

import (
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
)

type useCase struct {
	friendUseCase ports.FriendUseCase
}

func NewUseCase(friendUseCase ports.FriendUseCase) ports.UseCase {
	return &useCase{
		friendUseCase: friendUseCase,
	}
}

func (s *useCase) FriendUseCase() ports.FriendUseCase {
	return s.friendUseCase
}
