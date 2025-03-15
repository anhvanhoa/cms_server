package repository

import (
	"cms-server/internal/entity"

	"github.com/go-pg/pg/v10"
)

type MailProviderRepository interface {
	GetMailProviderByEmail(email string) (*entity.MailProvider, error)
}

type mailProviderRepositoryImpl struct {
	db *pg.DB
}

func NewMailProviderRepository(db *pg.DB) MailProviderRepository {
	return &mailProviderRepositoryImpl{
		db: db,
	}
}

func (mhr *mailProviderRepositoryImpl) GetMailProviderByEmail(email string) (*entity.MailProvider, error) {
	var mail entity.MailProvider
	err := mhr.db.Model(&mail).Where("email = ?", email).Select()
	return &mail, err
}
