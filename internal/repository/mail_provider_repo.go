package repository

import (
	"cms-server/internal/entity"
)

type MailProviderRepository interface {
	GetMailProviderByEmail(email string) (*entity.MailProvider, error)
}
