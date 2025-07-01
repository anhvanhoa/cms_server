package repo

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	"context"

	"github.com/go-pg/pg/v10"
)

type statusHistoryRepositoryImpl struct {
	db pg.DBI
}

func NewStatusHistoryRepository(db *pg.DB) repository.StatusHistoryRepository {
	return &statusHistoryRepositoryImpl{
		db: db,
	}
}

func (shr *statusHistoryRepositoryImpl) Create(data *entity.StatusHistory) error {
	_, err := shr.db.Model(data).Insert()
	return err
}

func (shr *statusHistoryRepositoryImpl) Tx(ctx context.Context) repository.StatusHistoryRepository {
	tx := getTx(ctx, shr.db)
	return &statusHistoryRepositoryImpl{
		db: tx,
	}
}
