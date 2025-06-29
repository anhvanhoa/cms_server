package repository

import (
	"cms-server/internal/entity"

	"github.com/go-pg/pg/v10"
)

type StatusHistoryRepository interface {
	Create(data *entity.StatusHistory, txs ...*pg.Tx) error
}
