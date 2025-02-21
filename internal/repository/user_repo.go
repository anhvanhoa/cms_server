package repository

import (
	"cms-server/internal/entity"

	"github.com/go-pg/pg/v10"
)

type UserRepository interface {
	CreateUser() error
	GetUserByEmailOrPhone(val string) (entity.User, error)
}

type userRepositoryImpl struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (ur *userRepositoryImpl) CreateUser() error {
	return nil
}

func (ur *userRepositoryImpl) GetUserByEmailOrPhone(val string) (entity.User, error) {
	var user entity.User
	err := ur.db.Model(&user).Where("email = ?", val).WhereOr("phone = ?", val).Select()
	return user, err
}
