package repository

import "cms-server/domain/entity"

type TypeMailRepo interface {
	Create(typeMail entity.TypeMail) (entity.TypeMail, error)
	GetByID(id string) (*entity.TypeMail, error)
	GetAll() ([]*entity.TypeMail, error)
	GetAllWithPagination(limit, offset int) ([]*entity.TypeMail, int, error)
	Update(id string, typeMail entity.TypeMail, cols ...string) error
	Delete(id string) error
}
