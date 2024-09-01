package repository

import (
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	platformPorts "github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

type repositories struct {
	friendRepo ports.FriendRepository
	outboxRepo ports.OutboxRepository
}

func NewRepositories(conn platformPorts.Connection) ports.Repository {
	return &repositories{
		friendRepo: NewFriendRepository(conn),
		outboxRepo: NewOutboxRepository(conn),
	}
}

func (r *repositories) Friend() ports.FriendRepository {
	return r.friendRepo
}

func (r *repositories) Outbox() ports.OutboxRepository {
	return r.outboxRepo
}
