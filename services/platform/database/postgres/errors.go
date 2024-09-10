package postgres

import "errors"

var (
	ErrNoRows              = errors.New("no rows in result set")
	ErrTxAlreadyStarted    = errors.New("transaction already started")
	ErrInvalidQueryBuilder = errors.New("invalid query builder")
)
