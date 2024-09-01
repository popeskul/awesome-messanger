package transaction

import (
	"context"
	"fmt"

	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/connection"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

type Manager struct {
	conn connection.Connection
}

func NewManager(conn connection.Connection) *Manager {
	return &Manager{conn: conn}
}

func (m *Manager) RunInTransaction(ctx context.Context, fn func(ctx context.Context, tx ports.Transaction) error) error {
	tx, err := m.conn.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(ctx, tx)
	return err
}
