package auth

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	pkgjwt "cms-server/pkg/jwt"
	"time"
)

type VerifyAccountUsecase interface {
	VerifyRegister(t string) (*pkgjwt.RegisterClaims, error)
	GetUserById(id string) (entity.User, error)
	VerifyAccount(id string) error
}

type verifyAccountUsecaseImpl struct {
	userRepo repository.UserRepository
	jwt      pkgjwt.JWT
}

func NewVerifyAccountUsecase(
	userRepo repository.UserRepository,
	jwt pkgjwt.JWT,
) VerifyAccountUsecase {
	return &verifyAccountUsecaseImpl{
		userRepo,
		jwt,
	}
}

func (u *verifyAccountUsecaseImpl) VerifyRegister(t string) (*pkgjwt.RegisterClaims, error) {
	data, err := u.jwt.VerifyRegisterToken(t)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *verifyAccountUsecaseImpl) GetUserById(id string) (entity.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *verifyAccountUsecaseImpl) VerifyAccount(id string) error {
	t := time.Now()
	user := entity.User{
		ID:         id,
		CodeVerify: "",
		Veryfied:   &t,
	}
	_, err := u.userRepo.UpdateUser(id, user)
	return err
}
