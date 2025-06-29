package repo

import (
	"cms-server/constants"
	"cms-server/internal/entity"
	"cms-server/internal/repository"

	"github.com/go-pg/pg/v10"
)
type mailTemplateRepositoryImpl struct {
	db *pg.DB
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
