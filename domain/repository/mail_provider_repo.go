package repository

import (
	"cms-server/domain/entity"
	"context"
)

type MailProviderRepository interface {
	GetMailProviderByEmail(email string) (*entity.MailProvider, error)
	Tx(ctx context.Context) MailProviderRepository
}
