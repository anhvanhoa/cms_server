package repository

import (
	"cms-server/internal/entity"
)

type MailTemplateRepository interface {
	GetMailTplById(id string) (*entity.MailTemplate, error)
}
