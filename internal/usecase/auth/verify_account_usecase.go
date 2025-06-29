package authUC

import (
	pkgjwt "cms-server/infrastructure/service/jwt"
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	serviceJwt "cms-server/internal/service/jwt"
	"time"
)

type VerifyAccountUsecase interface {
	VerifyRegister(t string) (*serviceJwt.VerifyClaims, error)
	GetUserById(id string) (entity.User, error)
	VerifyAccount(id string) error
}

type verifyAccountUsecaseImpl struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwt         serviceJwt.JwtService
}

func NewVerifyAccountUsecase(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwt serviceJwt.JwtService,
) VerifyAccountUsecase {
	return &verifyAccountUsecaseImpl{
		userRepo,
		sessionRepo,
		jwt,
	}
}

func (u *verifyAccountUsecaseImpl) VerifyRegister(t string) (*serviceJwt.VerifyClaims, error) {
	if isExist := u.sessionRepo.TokenExists(t); !isExist {
		return nil, pkgjwt.ErrTokenNotFound
	}
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
