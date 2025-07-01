package entity

import (
	"time"
)

type StatusOrder struct {
	tableName   struct{}  `pg:"status_orders,alias:so"`
	ID          string    `pg:"id,pk"`
	Name        string    `pg:"name"`
	Description string    `pg:"description"`
	Variant     string    `pg:"variant"`
	CreatedAt   time.Time `pg:"created_at"`
}

func (so *StatusOrder) NameTable() any {
	return so.tableName
}
