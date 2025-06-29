package repo

import (
	"cms-server/internal/repository"
	"context"

	"github.com/go-pg/pg/v10"
)

type managerTransaction struct {
	db *pg.DB
}

func NewManagerTransaction(db *pg.DB) repository.ManagerTransaction {
	return &managerTransaction{
		db: db,
	}
}

func (mt *managerTransaction) RunInTransaction(fn func(tx *pg.Tx) error) error {
	tx, err := mt.db.BeginContext(mt.db.Context())
	if err != nil {
		return err
	}

	// Nếu có lỗi thì rollback
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit()
}

type txContextKey struct{}

func (mt *managerTransaction) Do(fn func(ctx context.Context) error) error {
	tx, err := mt.db.BeginContext(mt.db.Context())
	if err != nil {
		return err
	}
	txCtx := context.WithValue(mt.db.Context(), txContextKey{}, tx)
	// Nếu có lỗi thì rollback
	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit()
}
