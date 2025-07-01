package entity

import (
	"time"
)

type ModuleChildRole struct {
	tableName struct{}   `pg:"module_child_roles,alias:mcr"`
	ID        string     `pg:"id,pk"`
	RoleID    string     `pg:"role_id"`
	ModuleID  string     `pg:"module_id"`
	ChildID   string     `pg:"child_id"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (mcr *ModuleChildRole) NameTable() any {
	return mcr.tableName
}
