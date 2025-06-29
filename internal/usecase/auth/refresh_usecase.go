package authUC

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	serviceJwt "cms-server/internal/service/jwt"
	"time"
)

type RefreshUsecase interface {
	GetSessionByToken(token string) (entity.Session, error)
	ClearSessionExpired() error
	VerifyToken(token string) (*serviceJwt.AuthClaims, error)
	GengerateAccessToken(id string, fullName string, exp time.Time) (string, error)
	GengerateRefreshToken(id string, fullName string, exp time.Time, os string) (string, error)
}

type refreshUsecaseImpl struct {
	sessionRepo repository.SessionRepository
	jwtAccess   serviceJwt.JwtService
	jwtRefresh  serviceJwt.JwtService
}

func NewRefreshUsecase(
	sessionRepo repository.SessionRepository,
	jwtAccess serviceJwt.JwtService,
	jwtRefresh serviceJwt.JwtService,
) RefreshUsecase {
	return &refreshUsecaseImpl{
		sessionRepo: sessionRepo,
		jwtAccess:   jwtAccess,
		jwtRefresh:  jwtRefresh,
	}
}

func (uc *refreshUsecaseImpl) GetSessionByToken(token string) (entity.Session, error) {
	session, err := uc.sessionRepo.GetSessionByToken(token)
	if err != nil {
		return entity.Session{}, err
	}
	if err := uc.sessionRepo.DeleteSessionAuthByToken(token); err != nil {
		return entity.Session{}, err
	}
	return session, nil
}

func (uc *refreshUsecaseImpl) ClearSessionExpired() error {
	if err := uc.sessionRepo.DeleteAllSessionsExpired(); err != nil {
		return err
	}
	return nil
}

func (uc *refreshUsecaseImpl) VerifyToken(token string) (*serviceJwt.AuthClaims, error) {
	claims, err := uc.jwtRefresh.VerifyAuthToken(token)
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func (uc *refreshUsecaseImpl) GengerateAccessToken(id string, fullName string, exp time.Time) (string, error) {
	return uc.jwtAccess.GenAuthToken(id, fullName, exp)
}

func (uc *refreshUsecaseImpl) GengerateRefreshToken(id string, fullName string, exp time.Time, os string) (string, error) {
	token, err := uc.jwtRefresh.GenAuthToken(id, fullName, exp)
	if err != nil {
		return "", err
	}
	if err := uc.sessionRepo.CreateSession(entity.Session{
		Token:     token,
		UserID:    id,
		Os:        os,
		Type:      entity.SessionTypeAuth,
		ExpiredAt: exp,
		CreatedAt: time.Now(),
	}); err != nil {
		return "", err
	}
	return token, nil
}
