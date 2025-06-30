package repository

import (
	"cms-server/internal/entity"
	"context"
)

type MailTemplateRepository interface {
	GetMailTplById(id string) (*entity.MailTemplate, error)
	Tx(ctx context.Context) MailTemplateRepository
}
