package typeMail

import (
	"cms-server/domain/common"
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	serviceError "cms-server/domain/service/error"
	"math"
)

var (
	ErrGetAllTypeMail = serviceError.NewErrorApp("Không thể lấy danh sách loại mail")
)

type GetAllUseCase interface {
	GetAll(limit, offset int) (common.PaginationResult[*entity.TypeMail], error)
}

type getAllUseCaseImpl struct {
	typeMailRepo repository.TypeMailRepo
}

func NewGetAllUseCase(
	typeMailRepo repository.TypeMailRepo,
) GetAllUseCase {
	return &getAllUseCaseImpl{
		typeMailRepo,
	}
}

func (uc *getAllUseCaseImpl) GetAll(limit, offset int) (common.PaginationResult[*entity.TypeMail], error) {
	typeMails, total, err := uc.typeMailRepo.GetAllWithPagination(limit, offset)
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	result := common.PaginationResult[*entity.TypeMail]{
		Data:       typeMails,
		Total:      total,
		TotalPages: totalPages,
		PageSize:   limit,
		Page:       totalPages + 1,
	}

	if err != nil {
		return result, ErrGetAllTypeMail
	}
	return result, nil
}
