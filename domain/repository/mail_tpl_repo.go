package repository

import (
	"cms-server/domain/entity"
	serviceError "cms-server/domain/service/error"
	"context"
)

var (
	ErrNotFoundTpl = serviceError.NewErrorApp("Không tìm thấy mẫu email")
)

type MailTemplateRepository interface {
	GetMailTplById(id string) (*entity.MailTemplate, error)
	Tx(ctx context.Context) MailTemplateRepository
}
