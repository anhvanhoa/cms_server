package repo

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"

	"github.com/go-pg/pg/v10"
)

type statusHistoryRepositoryImpl struct {
	db *pg.DB
}

func NewStatusHistoryRepository(db *pg.DB) repository.StatusHistoryRepository {
	return &statusHistoryRepositoryImpl{
		db: db,
	}
}

func (shr *statusHistoryRepositoryImpl) Create(data *entity.StatusHistory, txs ...*pg.Tx) error {
	if len(txs) > 0 {
		_, err := txs[0].Model(data).Insert()
		return err
	}
	_, err := shr.db.Model(data).Insert()
	return err
}
