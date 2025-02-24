package entity

import (
	"time"
)

type ActivityLog struct {
	tableName struct{}  `pg:"activity_logs,alias:al"`
	ID        string    `pg:"id,pk"`
	UserID    string    `pg:"user_id"`
	User      *User     `pg:"rel:has-one"`
	Url       string    `pg:"url"`
	Activity  string    `pg:"activity"`
	IP        string    `pg:"ip"`
	CreatedAt time.Time `pg:"created_at"`
}

func (al *ActivityLog) NameTable() any {
	return al.tableName
}
