package repository

import (
	"cms-server/domain/entity"
	"context"
)

type StatusHistoryRepository interface {
	Create(data *entity.StatusHistory) error
	Tx(ctx context.Context) StatusHistoryRepository
}
