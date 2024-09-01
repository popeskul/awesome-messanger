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

type OutboxRepository interface {
	Add(ctx context.Context, event models.OutboxEvent) error
	GetPendingEvents(ctx context.Context, limit int) ([]models.OutboxEvent, error)
	MarkAsProcessed(ctx context.Context, id int) error
}

type Repository interface {
	Friend() FriendRepository
	Outbox() OutboxRepository
}
