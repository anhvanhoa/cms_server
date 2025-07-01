package typeMail

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	serviceError "cms-server/domain/service/error"
	"time"
)

var (
	ErrUpdateTypeMail   = serviceError.NewErrorApp("Không thể cập nhật loại mail")
	ErrTypeMailNotFound = serviceError.NewErrorApp("Loại mail không tồn tại")
)

type UpdateUseCase interface {
	Update(id string, name string) error
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

func (uc *updateUseCaseImpl) Update(id string, name string) error {
	if _, err := uc.typeMailRepo.GetByID(id); err != nil {
		return ErrTypeMailNotFound
	}

	updatedAt := time.Now()
	typeMail := entity.TypeMail{
		Name:      name,
		UpdatedAt: &updatedAt,
	}
	err := uc.typeMailRepo.Update(id, typeMail, "name", "updated_at")
	if err != nil {
		return err
	}
	return nil
}
