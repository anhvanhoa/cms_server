package authUC

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	serviceJwt "cms-server/internal/service/jwt"
	"time"

	"github.com/alexedwards/argon2id"
)

type LoginUsecase interface {
	GetUserByEmailOrPhone(val string) (entity.User, error)
	CheckHashPassword(password, hash string) bool
	GengerateAccessToken(id string, fullName string, exp time.Time) (string, error)
	GengerateRefreshToken(id string, fullName string, exp time.Time, os string) (string, error)
}

type loginUsecaseImpl struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtAccess   serviceJwt.JwtService
	jwtRefresh  serviceJwt.JwtService
}

func NewLoginUsecase(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtAccess serviceJwt.JwtService,
	jwtRefresh serviceJwt.JwtService,
) LoginUsecase {
	return &loginUsecaseImpl{
		userRepo,
		sessionRepo,
		jwtAccess,
		jwtRefresh,
	}
}

func (uc *loginUsecaseImpl) GetUserByEmailOrPhone(val string) (entity.User, error) {
	return uc.userRepo.GetUserByEmailOrPhone(val)
}

func (uc *loginUsecaseImpl) CheckHashPassword(password, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false
	}
	return match
}

func (uc *loginUsecaseImpl) GengerateAccessToken(id string, fullName string, exp time.Time) (string, error) {
	return uc.jwtAccess.GenAuthToken(id, fullName, exp)
}

func (uc *loginUsecaseImpl) GengerateRefreshToken(id string, fullName string, exp time.Time, os string) (string, error) {
	token, err := uc.jwtRefresh.GenAuthToken(id, fullName, exp)
	if err != nil {
		return "", err
	}
	if err := uc.sessionRepo.CreateSession(entity.Session{
		Token:     token,
		UserID:    id,
		Os:        os,
		Type:      (entity.SessionTypeAuth),
		ExpiredAt: exp,
		CreatedAt: time.Now(),
	}); err != nil {
		return "", err
	}
	return token, nil
}
