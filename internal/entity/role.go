package entity

import (
	"time"
)

type Role struct {
	tableName   struct{}   `pg:"roles,alias:r"`
	Key         string     `pg:"key,pk"`
	Name        string     `pg:"name,unique"`
	Description string     `pg:"description"`
	Status      string     `pg:"status"`
	CreatedAt   time.Time  `pg:"created_at"`
	UpdatedAt   *time.Time `pg:"updated_at"`
}

func (r *Role) GetNameTable() any {
	return r.tableName
}
