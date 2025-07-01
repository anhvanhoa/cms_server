package repo

import (
	"cms-server/constants"
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	"context"

	"github.com/go-pg/pg/v10"
)

type mailTemplateRepositoryImpl struct {
	db pg.DBI
}

func NewMailTplRepository(db *pg.DB) repository.MailTemplateRepository {
	return &mailTemplateRepositoryImpl{
		db: db,
	}
}

func (mtr *mailTemplateRepositoryImpl) GetMailTplById(id string) (*entity.MailTemplate, error) {
	var tml entity.MailTemplate
	err := mtr.db.Model(&tml).Where("id = ?", id).
		Where("status = ?", constants.STATUS_ACTICE).
		Select()
	return &tml, err
}

func (mtr *mailTemplateRepositoryImpl) Tx(ctx context.Context) repository.MailTemplateRepository {
	tx := getTx(ctx, mtr.db)
	return &mailTemplateRepositoryImpl{
		db: tx,
	}
}
