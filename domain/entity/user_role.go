package entity

import (
	"time"
)

type UserRole struct {
	tableName struct{}   `pg:"user_roles,alias:ur"`
	ID        string     `pg:"id,pk"`
	UserID    string     `pg:"user_id,unique:role_key"`
	RoleKey   string     `pg:"role_key,unique:user_id"`
	CreatedBy string     `pg:"created_by"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (ur *UserRole) GetNameTable() any {
	return ur.tableName
}
