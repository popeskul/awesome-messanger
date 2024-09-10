package connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/config"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

type Connection struct {
	pool   *pgxpool.Pool
	config *config.Config
}

func New(ctx context.Context, cfg *config.Config) (ports.Connection, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConnections
	poolConfig.MinConns = cfg.MinConnections
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime
	poolConfig.HealthCheckPeriod = cfg.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return &Connection{
		pool:   pool,
		config: cfg,
	}, nil
}

func (c *Connection) Close() {
	c.pool.Close()
}

func (c *Connection) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return c.pool.Exec(ctx, query, args...)
}

func (c *Connection) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return c.pool.Query(ctx, query, args...)
}

func (c *Connection) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return c.pool.QueryRow(ctx, query, args...)
}

func (c *Connection) BeginTx(ctx context.Context) (ports.Transaction, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to begin transaction: %w", err)
	}
	return &Transaction{tx: tx}, nil
}

type Transaction struct {
	tx pgx.Tx
}

func (t *Transaction) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *Transaction) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

func (t *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return t.tx.Exec(ctx, query, args...)
}

func (t *Transaction) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return t.tx.Query(ctx, query, args...)
}

func (t *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return t.tx.QueryRow(ctx, query, args...)
}

func (c *Connection) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}
