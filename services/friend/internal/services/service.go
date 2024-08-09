package services

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/friend/internal/models"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Service struct {
	logger Logger
}

func NewService(logger Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) AddFriend(ctx context.Context, inout *models.Friend) (*models.Friend, error) {
	s.logger.Printf("Add friend request received for user %s", inout.UserId)

	return inout, nil
}

func (s *Service) GetFriends(ctx context.Context) ([]*models.Friend, error) {
	s.logger.Printf("Get friends request received for user %s", "1")

	return []*models.Friend{
		{
			UserId:   "1",
			FriendId: "2",
		},
	}, nil
}

func (s *Service) RespondToFriendRequest(ctx context.Context, inout *models.Friend) (*models.Friend, error) {
	s.logger.Printf("Respond friend request received for user %s", inout.UserId)

	return inout, nil
}
