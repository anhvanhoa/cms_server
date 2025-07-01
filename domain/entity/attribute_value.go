package entity

import (
	"time"
)

type AttributeValue struct {
	tableName   struct{}   `pg:"attribute_values,alias:av"`
	ID          string     `pg:"id,pk"`
	AttributeID string     `pg:"attribute_id"`
	Attribute   *Attribute `pg:"rel:has-one"`
	Value       string     `pg:"value"`
	CreatedAt   time.Time  `pg:"created_at"`
	UpdatedAt   *time.Time `pg:"updated_at"`
}

func (a *AttributeValue) NameTable() any {
	return a.tableName
}
