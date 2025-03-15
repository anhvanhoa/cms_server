package repository

import "github.com/go-pg/pg/v10"

type ManagerTransaction interface {
	RunInTransaction(fn func(tx *pg.Tx) error) error
}

type managerTransaction struct {
	db *pg.DB
}

func NewManagerTransaction(db *pg.DB) *managerTransaction {
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
