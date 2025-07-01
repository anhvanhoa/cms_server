package repo

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"

	"github.com/go-pg/pg/v10"
)

type typeMailRepositoryImpl struct {
	db pg.DBI
}

func NewTypeMailRepository(db *pg.DB) repository.TypeMailRepo {
	return &typeMailRepositoryImpl{
		db: db,
	}
}

func (tmr *typeMailRepositoryImpl) Create(typeMail entity.TypeMail) (entity.TypeMail, error) {
	_, err := tmr.db.Model(&typeMail).Insert()
	return typeMail, err
}

func (tmr *typeMailRepositoryImpl) GetByID(id string) (*entity.TypeMail, error) {
	var tm entity.TypeMail
	err := tmr.db.Model(&tm).Where("id = ?", id).Select()
	return &tm, err
}

func (tmr *typeMailRepositoryImpl) GetAllWithPagination(limit, offset int) ([]*entity.TypeMail, int, error) {
	var tms []*entity.TypeMail
	total, err := tmr.db.Model(&tms).Limit(limit).Offset(offset).SelectAndCount()
	return tms, total, err
}

func (tmr *typeMailRepositoryImpl) GetAll() ([]*entity.TypeMail, error) {
	var tms []*entity.TypeMail
	err := tmr.db.Model(&tms).Select()
	return tms, err
}

func (tmr *typeMailRepositoryImpl) Update(typeMail entity.TypeMail) error {
	_, err := tmr.db.Model(&typeMail).WherePK().Update()
	return err
}

func (tmr *typeMailRepositoryImpl) Delete(id string) error {
	_, err := tmr.db.Model(&entity.TypeMail{}).Where("id = ?", id).Delete()
	return err
}
