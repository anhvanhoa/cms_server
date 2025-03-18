package repository

import (
	"cms-server/internal/entity"
	"fmt"
	"reflect"
	"strings"

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

func (ur *userRepositoryImpl) UpdateUser(id string, user entity.User, txs ...*pg.Tx) (entity.UserInfor, error) {
	var setClauses []string
	var params []interface{}

	// Sử dụng reflection để lấy danh sách field cần cập nhật
	v := reflect.ValueOf(user)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Lấy giá trị thực nếu là pointer
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Bỏ qua các trường không cần cập nhật
		if field.Name == "ID" || field.Name == "CreatedAt" {
			continue
		}

		// Chỉ thêm vào danh sách cập nhật nếu giá trị không phải zero-value
		if !value.IsZero() {
			columnName := field.Tag.Get("pg")
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", columnName))
			params = append(params, value.Interface())
		}
	}

	// Nếu không có dữ liệu để cập nhật, return sớm
	if len(setClauses) == 0 {
		return user.GetInfor(), nil
	}

	setQuery := strings.Join(setClauses, ", ")

	var err error
	if len(txs) > 0 {
		_, err = txs[0].Model(&user).Where("id = ?", id).Set(setQuery, params...).Update()
	} else {
		_, err = ur.db.Model(&user).Where("id = ?", id).Set(setQuery, params...).Update()
	}

	return user.GetInfor(), err
}

func (ur *userRepositoryImpl) GetUserByID(id string) (entity.User, error) {
	var user entity.User
	err := ur.db.Model(&user).Where("id = ?", id).Select()
	return user, err
}

func (ur *userRepositoryImpl) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := ur.db.Model(&user).Where("email = ?", email).Select()
	return user, err
}

func (ur *userRepositoryImpl) UpdateUserByEmail(email string, user entity.User, txs ...*pg.Tx) (bool, error) {
	if len(txs) > 0 {
		r, err := txs[0].Model(&user).Where("email = ?", email).Update()
		return r.RowsAffected() != -1, err
	}
	r, err := ur.db.Model(&user).Where("email = ?", email).Update(&user)
	return r.RowsAffected() != -1, err
}
