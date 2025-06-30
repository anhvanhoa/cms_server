package repository

import (
	"cms-server/internal/entity"
	"context"
)

type MailHistoryRepository interface {
	Create(data *entity.MailHistory) error
	UpdateSubAndBodyById(id, sub, body string) error
	GetMailHistoryById(id string) (*entity.MailHistory, error)
	Tx(ctx context.Context) MailHistoryRepository
}
