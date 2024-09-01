package postgres

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/config"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/connection"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

func NewDatabase(ctx context.Context, cfg *config.Config) (ports.Connection, error) {
	return connection.New(ctx, cfg)
}
