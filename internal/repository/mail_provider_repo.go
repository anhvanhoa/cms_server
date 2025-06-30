package repository

import (
	"cms-server/internal/entity"
	"context"
)

type MailProviderRepository interface {
	GetMailProviderByEmail(email string) (*entity.MailProvider, error)
	Tx(ctx context.Context) MailProviderRepository
}
