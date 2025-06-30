package repo

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	"context"

	"github.com/go-pg/pg/v10"
)

type mailHistoryRepositoryImpl struct {
	db pg.DBI
}

func NewMailHistoryRepository(db *pg.DB) repository.MailHistoryRepository {
	return &mailHistoryRepositoryImpl{
		db: db,
	}
}

func (mhr *mailHistoryRepositoryImpl) Create(data *entity.MailHistory) error {
	_, err := mhr.db.Model(data).Insert()
	return err
}

func (mhr *mailHistoryRepositoryImpl) UpdateSubAndBodyById(id, sub, body string) error {
	var m entity.MailHistory
	_, err := mhr.db.Model(&m).Where("id = ?", id).Set("subject = ?", sub).Set("body = ?", body).Update()
	return err
}

func (mhr *mailHistoryRepositoryImpl) GetMailHistoryById(id string) (*entity.MailHistory, error) {
	var mail entity.MailHistory
	err := mhr.db.Model(&mail).Where("id = ?", id).Select()
	return &mail, err
}

func (mhr *mailHistoryRepositoryImpl) Tx(ctx context.Context) repository.MailHistoryRepository {
	tx := getTx(ctx, mhr.db)
	return &mailHistoryRepositoryImpl{
		db: tx,
	}
}
