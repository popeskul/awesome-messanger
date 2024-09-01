package usecase

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
)

type friendUseCase struct {
	logger ports.Logger
	repo   ports.FriendRepository
}

func NewFriendUseCase(logger ports.Logger, repo ports.FriendRepository) ports.FriendUseCase {
	return &friendUseCase{
		logger: logger,
		repo:   repo,
	}
}

func (s *friendUseCase) AddFriend(ctx context.Context, inout *models.Friend) (*models.Friend, error) {
	s.logger.Info("Add friend request received for user %s", inout.UserId)

	friend, err := s.repo.AddFriend(ctx, inout)
	if err != nil {
		s.logger.Error("Error adding friend: %v", err)
		return nil, err
	}

	return friend, nil
}

func (s *friendUseCase) GetFriends(ctx context.Context) ([]*models.Friend, error) {
	s.logger.Info("Get friends request received for user %s", "1")

	friends, err := s.repo.GetFriends(ctx, "1")
	if err != nil {
		s.logger.Error("Error getting friends: %v", err)
		return nil, err
	}

	return friends, nil
}

func (s *friendUseCase) RespondToFriendRequest(ctx context.Context, inout *models.Friend) (*models.Friend, error) {
	s.logger.Info("Respond friend request received for user %s", inout.UserId)
	return inout, nil
}
