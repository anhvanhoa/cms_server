package entity

import (
	"time"
)

type Attribute struct {
	tableName struct{}   `pg:"attributes,alias:a"`
	ID        string     `pg:"id,pk"`
	Name      string     `pg:"name"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (a *Attribute) NameTable() any {
	return a.tableName
}
