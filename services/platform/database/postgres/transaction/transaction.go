package transaction

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Transaction struct {
	tx pgx.Tx
}

func New(tx pgx.Tx) *Transaction {
	return &Transaction{tx: tx}
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
