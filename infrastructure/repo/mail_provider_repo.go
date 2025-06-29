package repo

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"

	"github.com/go-pg/pg/v10"
)

type mailProviderRepositoryImpl struct {
	db *pg.DB
}

func NewMailProviderRepository(db *pg.DB) repository.MailProviderRepository {
	return &mailProviderRepositoryImpl{
		db: db,
	}
}

func (mhr *mailProviderRepositoryImpl) GetMailProviderByEmail(email string) (*entity.MailProvider, error) {
	var mail entity.MailProvider
	err := mhr.db.Model(&mail).Where("email = ?", email).Select()
	return &mail, err
}
