package typeMail

import (
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	serviceError "cms-server/domain/service/error"
	"cms-server/domain/service/goid"
	"time"
)

var (
	ErrCreateTypeMail = serviceError.NewErrorApp("Không thể tạo loại mail")
)

type CreateUseCase interface {
	Create(name string, createdBy string) error
}

type createUseCaseImpl struct {
	typeMailRepo repository.TypeMailRepo
	goid         goid.GoId
}

func NewCreateUseCase(
	typeMailRepo repository.TypeMailRepo,
	goid goid.GoId,
) CreateUseCase {
	return &createUseCaseImpl{
		typeMailRepo,
		goid,
	}
}

func (uc *createUseCaseImpl) Create(name, createdBy string) error {
	Id := uc.goid.GenWithLength(7)
	typeMail := entity.TypeMail{
		ID:        Id,
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
	_, err := uc.typeMailRepo.Create(typeMail)
	if err != nil {
		return ErrCreateTypeMail
	}
	return nil
}
