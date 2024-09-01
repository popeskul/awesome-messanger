package repository

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	platformPorst "github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

type friendRepository struct {
	conn platformPorst.Connection
}

func NewFriendRepository(conn platformPorst.Connection) ports.FriendRepository {
	return &friendRepository{conn: conn}
}

func (r *friendRepository) AddFriend(ctx context.Context, friend *models.Friend) (*models.Friend, error) {
	query := `INSERT INTO friends (user_id, friend_id) VALUES ($1, $2) RETURNING user_id, friend_id`
	var newFriend models.Friend
	err := r.conn.QueryRow(ctx, query, friend.UserId, friend.FriendId).Scan(&newFriend.UserId, &newFriend.FriendId)
	if err != nil {
		return nil, err
	}
	return &newFriend, nil
}

func (r *friendRepository) GetFriends(ctx context.Context, userID string) ([]*models.Friend, error) {
	query := `SELECT user_id, friend_id FROM friends WHERE user_id = $1`
	rows, err := r.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []*models.Friend
	for rows.Next() {
		var friend models.Friend
		if err := rows.Scan(&friend.UserId, &friend.FriendId); err != nil {
			return nil, err
		}
		friends = append(friends, &friend)
	}
	return friends, nil
}

func (r *friendRepository) RemoveFriend(ctx context.Context, userID, friendID string) error {
	query := `DELETE FROM friends WHERE user_id = $1 AND friend_id = $2`
	_, err := r.conn.Exec(ctx, query, userID, friendID)
	return err
}
