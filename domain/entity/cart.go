package entity

import (
	"time"
)

type Cart struct {
	tableName     struct{}  `pg:"carts,alias:c"`
	ID            string    `pg:"id,pk"`
	UserID        string    `pg:"user_id"`
	User          *User     `pg:"rel:has-one"`
	SessionsToken string    `pg:"sessions_token"`
	Sessions      *Session  `pg:"rel:has-one"`
	CreatedAt     time.Time `pg:"created_at"`
}

func (c *Cart) NameTable() any {
	return c.tableName
}
