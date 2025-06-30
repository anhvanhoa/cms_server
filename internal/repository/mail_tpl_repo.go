package repository

import (
	"cms-server/internal/entity"
	serviceError "cms-server/internal/service/error"
	"context"
)

var (
	ErrNotFoundTpl = serviceError.NewErrorApp("Không tìm thấy mẫu email")
)

type MailTemplateRepository interface {
	GetMailTplById(id string) (*entity.MailTemplate, error)
	Tx(ctx context.Context) MailTemplateRepository
}
