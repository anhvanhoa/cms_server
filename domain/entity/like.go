package entity

import (
	"time"
)

type Like struct {
	tableName struct{}  `pg:"likes,alias:l"`
	ID        string    `pg:"id,pk"`
	UserID    string    `pg:"user_id"`
	User      *User     `pg:"rel:has-one"`
	RefId     string    `pg:"ref_id"`
	Type      string    `pg:"type"`
	CreatedAt time.Time `pg:"created_at"`
}

func (l *Like) NameTable() any {
	return l.tableName
}
