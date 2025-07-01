package typeMail

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	serviceError "cms-server/domain/service/error"
	"time"
)

var (
	ErrUpdateTypeMail = serviceError.NewErrorApp("Không thể cập nhật loại mail")
)

type UpdateUseCase interface {
	Update(id string, name string, updatedBy string) error
}

type updateUseCaseImpl struct {
	typeMailRepo repository.TypeMailRepo
}

func NewUpdateUseCase(
	typeMailRepo repository.TypeMailRepo,
) UpdateUseCase {
	return &updateUseCaseImpl{
		typeMailRepo,
	}
}

func (uc *updateUseCaseImpl) Update(id string, name string, updatedBy string) error {
	updatedAt := time.Now()
	typeMail := entity.TypeMail{
		Name:      name,
		UpdatedAt: &updatedAt,
	}
	err := uc.typeMailRepo.Update(typeMail)
	if err != nil {
		return ErrUpdateTypeMail
	}
	return nil
}
