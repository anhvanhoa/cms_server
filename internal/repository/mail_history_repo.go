package repository

import (
	"cms-server/internal/entity"

	"github.com/go-pg/pg/v10"
)

type MailHistoryRepository interface {
	Create(data *entity.MailHistory, txs ...*pg.Tx) error
	UpdateSubAndBodyById(id, sub, body string, txs ...*pg.Tx) error
	GetMailHistoryById(id string) (*entity.MailHistory, error)
}
