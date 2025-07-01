package typeMail

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"
)

type GetUseCase interface {
	GetByID(id string) (*entity.TypeMail, error)
}

type getUseCaseImpl struct {
	typeMailRepo repository.TypeMailRepo
}

func NewGetUseCase(
	typeMailRepo repository.TypeMailRepo,
) GetUseCase {
	return &getUseCaseImpl{
		typeMailRepo,
	}
}

func (uc *getUseCaseImpl) GetByID(id string) (*entity.TypeMail, error) {
	typeMail, err := uc.typeMailRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return typeMail, nil
}
