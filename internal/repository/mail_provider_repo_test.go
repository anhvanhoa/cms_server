package repository

import (
	"cms-server/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockMailProviderRepo struct{}

func (m *mockMailProviderRepo) GetMailProviderByEmail(email string) (*entity.MailProvider, error) {
	if email == "found@example.com" {
		return &entity.MailProvider{Email: email, Name: "Provider"}, nil
	}
	return nil, nil
}
func (m *mockMailProviderRepo) Tx(ctx context.Context) MailProviderRepository { return m }

func TestMailProviderRepo_GetMailProviderByEmail(t *testing.T) {
	repo := &mockMailProviderRepo{}
	provider, err := repo.GetMailProviderByEmail("found@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, provider)
	assert.Equal(t, "found@example.com", provider.Email)

	none, err := repo.GetMailProviderByEmail("notfound@example.com")
	assert.NoError(t, err)
	assert.Nil(t, none)
}
