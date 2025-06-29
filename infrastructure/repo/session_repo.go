package repo

import (
	"cms-server/internal/entity"
	"cms-server/internal/repository"

	"github.com/go-pg/pg/v10"
)

type sessionRepositoryImpl struct {
	db *pg.DB
}

func NewSessionRepository(db *pg.DB) repository.SessionRepository {
	return &sessionRepositoryImpl{
		db: db,
	}
}

func (sr *sessionRepositoryImpl) CreateSession(data entity.Session) error {
	_, err := sr.db.Model(&data).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (sr *sessionRepositoryImpl) GetSessionByToken(token string) (entity.Session, error) {
	var session entity.Session
	err := sr.db.Model(&session).Where("token = ?", token).Select()
	if err != nil {
		return session, err
	}
	return session, nil
}

func (sr *sessionRepositoryImpl) TokenExists(token string) bool {
	count, err := sr.db.Model(&entity.Session{}).Where("token = ?", token).Count()
	if err != nil {
		return false
	}
	return count > 0
}

func (sr *sessionRepositoryImpl) DeleteSessionVerifyByUserID(userID string) error {
	return sr.DeleteSessionByTypeAndUserID(entity.SessionTypeVerify, userID)
}

func (sr *sessionRepositoryImpl) DeleteSessionByTypeAndUserID(sessionType entity.SessionType, userID string) error {
	_, err := sr.db.Model(&entity.Session{}).
		Where("type = ? AND user_id = ?", sessionType, userID).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

func (sr *sessionRepositoryImpl) DeleteSessionByTypeAndToken(sessionType entity.SessionType, token string) error {
	_, err := sr.db.Model(&entity.Session{}).
		Where("type = ? AND token = ?", sessionType, token).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

func (sr *sessionRepositoryImpl) DeleteSessionAuthByToken(token string) error {
	return sr.DeleteSessionByTypeAndToken(entity.SessionTypeAuth, token)
}

func (sr *sessionRepositoryImpl) DeleteAllSessionsExpired() error {
	_, err := sr.db.Model(&entity.Session{}).
		Where("expired_at < NOW()").
		Delete()
	if err != nil {
		return err
	}
	return nil
}
