package entity

import (
	"time"
)

type VariantValue struct {
	tableName        struct{}  `pg:"variant_values,alias:vv"`
	AttributeValueID string    `pg:"attribute_value_id,pk"`
	VariantID        string    `pg:"variant_id,pk"`
	CreatedAt        time.Time `pg:"created_at"`
}

func (vv *VariantValue) NameTable() any {
	return vv.tableName
}
