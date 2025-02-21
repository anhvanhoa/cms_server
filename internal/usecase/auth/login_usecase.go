package auth

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"
)

type LoginUsecase interface {
	GetUserByEmailOrPhone(val string) (entity.User, error)
}

type loginUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewLoginUsecase(userRepo repository.UserRepository) LoginUsecase {
	return &loginUsecaseImpl{
		userRepo: userRepo,
	}
}

func (uc *loginUsecaseImpl) GetUserByEmailOrPhone(val string) (entity.User, error) {
	return uc.userRepo.GetUserByEmailOrPhone(val)
}
