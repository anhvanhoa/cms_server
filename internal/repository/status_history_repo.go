package repository

import (
	"cms-server/internal/entity"
	"context"
)

type StatusHistoryRepository interface {
	Create(data *entity.StatusHistory) error
	Tx(ctx context.Context) StatusHistoryRepository
}
