package repo

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	"context"

	"github.com/go-pg/pg/v10"
)

type mailProviderRepositoryImpl struct {
	db pg.DBI
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

func (mhr *mailProviderRepositoryImpl) Tx(ctx context.Context) repository.MailProviderRepository {
	tx := getTx(ctx, mhr.db)
	return &mailProviderRepositoryImpl{
		db: tx,
	}
}
