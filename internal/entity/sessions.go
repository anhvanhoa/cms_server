package entity

import (
	"time"
)

type Session struct {
	tableName struct{}  `pg:"sessions,alias:s"`
	Token     string    `pg:"token,pk"`
	UserID    string    `pg:"user_id"`
	User      *User     `pg:"rel:has-one"`
	Type      string    `pg:"type"`
	Os        string    `pg:"os"`
	ExpiredAt time.Time `pg:"expired_at"`
	CreatedAt time.Time `pg:"created_at"`
}

func (s *Session) NameTable() any {
	return s.tableName
}
