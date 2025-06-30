package repository

import (
	"cms-server/internal/entity"
	"context"
)

type SessionRepository interface {
	CreateSession(data entity.Session) error
	GetSessionByToken(token string) (entity.Session, error)
	TokenExists(token string) bool
	DeleteSessionByTypeAndUserID(sessionType entity.SessionType, userID string) error
	DeleteSessionByTypeAndToken(sessionType entity.SessionType, token string) error
	DeleteSessionVerifyByUserID(userID string) error
	DeleteSessionAuthByToken(token string) error
	DeleteSessionVerifyByToken(token string) error
	DeleteAllSessionsExpired() error
	Tx(ctx context.Context) SessionRepository
}
