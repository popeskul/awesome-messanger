package repository

import (
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	platformPorst "github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

type repositories struct {
	friendRepo ports.FriendRepository
}

func NewRepositories(conn platformPorst.Connection) ports.Repository {
	return &repositories{
		friendRepo: NewFriendRepository(conn),
	}
}

func (r *repositories) Friend() ports.FriendRepository {
	return r.friendRepo
}
