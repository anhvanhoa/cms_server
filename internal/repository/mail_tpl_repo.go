package repository

import (
	"cms-server/internal/entity"

	"github.com/go-pg/pg/v10"
)

type MailTemplateRepository interface {
	GetMailTplById(id string) (*entity.MailTemplate, error)
}

type mailTemplateRepositoryImpl struct {
	db *pg.DB
}

func NewMailTplRepository(db *pg.DB) MailTemplateRepository {
	return &mailTemplateRepositoryImpl{
		db: db,
	}
}

func (mtr *mailTemplateRepositoryImpl) GetMailTplById(id string) (*entity.MailTemplate, error) {
	var tml entity.MailTemplate
	err := mtr.db.Model(&tml).Where("id = ?", id).Select()
	return &tml, err
}
