package ports

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
)

type FriendUseCase interface {
	GetFriends(ctx context.Context) ([]*models.Friend, error)
	AddFriend(ctx context.Context, inout *models.Friend) (*models.Friend, error)
	RespondToFriendRequest(ctx context.Context, inout *models.Friend) (*models.Friend, error)
}

type UseCase interface {
	FriendUseCase() FriendUseCase
}
