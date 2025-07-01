package repository

import (
	"cms-server/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct{}

func (m *mockUserRepo) CreateUser(entity.User) (entity.UserInfor, error) {
	return entity.UserInfor{ID: "1"}, nil
}
func (m *mockUserRepo) GetUserByEmailOrPhone(val string) (entity.User, error) {
	return entity.User{ID: "1", Email: val}, nil
}
func (m *mockUserRepo) GetUserByID(id string) (entity.User, error) { return entity.User{ID: id}, nil }
func (m *mockUserRepo) CheckUserExist(val string) (bool, error) {
	return val == "exist@example.com", nil
}
func (m *mockUserRepo) GetUserByEmail(email string) (entity.User, error) {
	return entity.User{ID: "1", Email: email}, nil
}
func (m *mockUserRepo) UpdateUser(Id string, data entity.User) (entity.UserInfor, error) {
	return entity.UserInfor{ID: Id}, nil
}
func (m *mockUserRepo) UpdateUserByEmail(email string, data entity.User) (bool, error) {
	return true, nil
}
func (m *mockUserRepo) Tx(ctx context.Context) UserRepository { return m }

func TestUserRepo_GetUserByEmailOrPhone(t *testing.T) {
	repo := &mockUserRepo{}
	user, err := repo.GetUserByEmailOrPhone("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestUserRepo_CheckUserExist(t *testing.T) {
	repo := &mockUserRepo{}
	exist, err := repo.CheckUserExist("exist@example.com")
	assert.NoError(t, err)
	assert.True(t, exist)
	notExist, err := repo.CheckUserExist("notfound@example.com")
	assert.NoError(t, err)
	assert.False(t, notExist)
}
