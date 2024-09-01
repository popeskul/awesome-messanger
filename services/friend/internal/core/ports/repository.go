package ports

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
)

type FriendRepository interface {
	AddFriend(ctx context.Context, friend *models.Friend) (*models.Friend, error)
	GetFriends(ctx context.Context, userID string) ([]*models.Friend, error)
	RemoveFriend(ctx context.Context, userID, friendID string) error
}

type Repository interface {
	Friend() FriendRepository
}
