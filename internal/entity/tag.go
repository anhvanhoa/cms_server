package entity

import (
	"time"
)

type Tag struct {
	tableName struct{}   `pg:"tags,alias:t"`
	ID        string     `pg:"id,pk"`
	Name      string     `pg:"name"`
	Variant   string     `pg:"variant"`
	Status    string     `pg:"status"`
	CreatedBy string     `pg:"created_by"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (t *Tag) NameTable() any {
	return t.tableName
}
