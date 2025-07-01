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
	GetAll(pageSize, page int) (common.PaginationResult[*entity.TypeMail], error)
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

func (uc *getAllUseCaseImpl) GetAll(pageSize, page int) (common.PaginationResult[*entity.TypeMail], error) {
	if pageSize <= 0 {
		pageSize = 10
	}
	if page <= 0 {
		page = 1
	}
	limit := pageSize
	offset := (page - 1) * pageSize
	typeMails, total, err := uc.typeMailRepo.GetAllWithPagination(limit, offset)
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	result := common.PaginationResult[*entity.TypeMail]{
		Data:       typeMails,
		Total:      total,
		TotalPages: totalPages,
		PageSize:   limit,
		Page:       page,
	}

	if err != nil {
		return result, ErrGetAllTypeMail
	}
	return result, nil
}
