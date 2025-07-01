package typeMail

import (
	"cms-server/domain/repository"
	serviceError "cms-server/domain/service/error"
)

var (
	ErrDeleteTypeMail = serviceError.NewErrorApp("Không thể xóa loại mail")
)

type DeleteUseCase interface {
	Delete(id string) error
}

type deleteUseCaseImpl struct {
	typeMailRepo repository.TypeMailRepo
}

func NewDeleteUseCase(
	typeMailRepo repository.TypeMailRepo,
) DeleteUseCase {
	return &deleteUseCaseImpl{
		typeMailRepo,
	}
}

func (uc *deleteUseCaseImpl) Delete(id string) error {
	err := uc.typeMailRepo.Delete(id)
	if err != nil {
		return ErrDeleteTypeMail
	}
	return nil
}
