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

type mailHistoryRepositoryImpl struct {
	db *pg.DB
}

func NewMailHistoryRepository(db *pg.DB) MailHistoryRepository {
	return &mailHistoryRepositoryImpl{
		db: db,
	}
}

func (mhr *mailHistoryRepositoryImpl) Create(data *entity.MailHistory, txs ...*pg.Tx) error {
	if len(txs) > 0 {
		_, err := txs[0].Model(data).Insert()
		return err
	}
	_, err := mhr.db.Model(data).Insert()
	return err
}

func (mhr *mailHistoryRepositoryImpl) UpdateSubAndBodyById(id, sub, body string, txs ...*pg.Tx) error {
	var m entity.MailHistory
	if len(txs) > 0 {
		_, err := txs[0].Model(&m).Where("id = ?", id).Set("subject = ?", sub).Set("body = ?", body).Update()
		return err
	}
	_, err := mhr.db.Model(&m).Where("id = ?", id).Set("subject = ?", sub).Set("body = ?", body).Update()
	return err
}

func (mhr *mailHistoryRepositoryImpl) GetMailHistoryById(id string) (*entity.MailHistory, error) {
	var mail entity.MailHistory
	err := mhr.db.Model(&mail).Where("id = ?", id).Select()
	return &mail, err
}
