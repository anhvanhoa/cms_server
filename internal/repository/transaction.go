package repository

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Tx interface {
	Commit() error
	Rollback() error
}

type ManagerTransaction interface {
	RunInTransaction(fn func(tx *pg.Tx) error) error
	Do(fn func(ctx context.Context) error) error
}
