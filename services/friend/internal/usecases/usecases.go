package usecases

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
)

type useCases struct {
	logger ports.Logger
}

func NewUseCases(logger ports.Logger) ports.UseCases {
	return &useCases{
		logger: logger,
	}
}

func (s *useCases) AddFriend(ctx context.Context, inout *models.Friend) (*models.Friend, error) {
	s.logger.Info("Add friend request received for user %s", inout.UserId)
	return inout, nil
}

func (s *useCases) GetFriends(ctx context.Context) ([]*models.Friend, error) {
	s.logger.Info("Get friends request received for user %s", "1")
	return []*models.Friend{
		{
			UserId:   "1",
			FriendId: "2",
		},
	}, nil
}

func (s *useCases) RespondToFriendRequest(ctx context.Context, inout *models.Friend) (*models.Friend, error) {
	s.logger.Info("Respond friend request received for user %s", inout.UserId)
	return inout, nil
}
