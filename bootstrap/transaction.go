package bootstrap

import (
	pkglog "cms-server/pkg/logger"

	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type TransactionManager struct {
	db  *pg.DB
	log pkglog.Logger
}

func NewTransactionManager(db *pg.DB, log pkglog.Logger) *TransactionManager {
	return &TransactionManager{db, log}
}

// func (t *TransactionManager) Begin() *pg.Tx {
// 	tx, err := t.db.Begin()
// 	if err != nil {
// 		t.log.Fatal("Error starting transaction", zap.String("error", err.Error()))
// 	}
// 	return tx
// }

// func (t *TransactionManager) Commit(tx *pg.Tx) error {
// 	return tx.Commit()
// }

// func (t *TransactionManager) Rollback(tx *pg.Tx) error {
// 	return tx.Rollback()
// }

func (tm *TransactionManager) WithTransaction(fn func(tx *pg.Tx) error) error {
	tx,err := tm.db.Begin() // Bắt đầu transaction
	if err != nil {
		tm.log.Fatal("Error starting transaction", zap.String("error", err.Error()))
	}
	if err := fn(tx); err != nil {
		tx.Rollback() // Rollback nếu có lỗi
		return err
	}
	return tx.Commit() // Commit nếu không có lỗi
}
