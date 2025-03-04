package repository

import (
	"cms-server/internal/entity"

	"github.com/go-pg/pg/v10"
)

type UserRepository interface {
	CreateUser(entity.User, ...*pg.Tx) (entity.UserInfor, error)
	GetUserByEmailOrPhone(val string) (entity.User, error)
	CheckUserExist(val string) (bool, error)
	RunInTransaction(fn func(tx *pg.Tx) error) error
}

type userRepositoryImpl struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (ur *userRepositoryImpl) RunInTransaction(fn func(tx *pg.Tx) error) error {
	tx, err := ur.db.BeginContext(ur.db.Context())
	if err != nil {
		return err
	}

	// Nếu có lỗi thì rollback
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit()
}

func (ur *userRepositoryImpl) CreateUser(user entity.User, txs ...*pg.Tx) (entity.UserInfor, error) {
	if len(txs) > 0 {
		_, err := txs[0].Model(&user).Insert()
		return user.GetInfor(), err
	}
	_, err := ur.db.Model(&user).Insert()
	return user.GetInfor(), err
}

func (ur *userRepositoryImpl) GetUserByEmailOrPhone(val string) (entity.User, error) {
	var user entity.User
	err := ur.db.Model(&user).Where("email = ?", val).WhereOr("phone = ?", val).Select()
	return user, err
}

func (ur *userRepositoryImpl) CheckUserExist(val string) (bool, error) {
	var user entity.User
	count, err := ur.db.Model(&user).Where("email = ?", val).Count()
	isExist := count > 0
	return isExist, err
}
