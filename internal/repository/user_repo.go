package repository

import (
	"cms-server/internal/entity"
	"context"

	"github.com/go-pg/pg/v10"
)

type UserRepository interface {
	CreateUser(entity.User, ...*pg.Tx) (entity.UserInfor, error)
	GetUserByEmailOrPhone(val string) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	CheckUserExist(val string) (bool, error)
	GetUserByEmail(email string) (entity.User, error)
	UpdateUser(Id string, data entity.User, txs ...*pg.Tx) (entity.UserInfor, error)
	UpdateUserByEmail(email string, data entity.User, txs ...*pg.Tx) (bool, error)
	Tx(ctx context.Context) UserRepository
}
