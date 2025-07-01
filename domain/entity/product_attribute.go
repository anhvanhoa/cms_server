package entity

import (
	"time"
)

type ProductAttribute struct {
	tableName   struct{}   `pg:"product_attributes,alias:pa"`
	ProductId   string     `pg:"product_id"`
	Product     *Product   `pg:"rel:has-one"`
	AttributeId string     `pg:"attribute_id"`
	Attribute   *Attribute `pg:"rel:has-one"`
	CreatedAt   time.Time  `pg:"created_at"`
}

func (a *ProductAttribute) NameTable() any {
	return a.tableName
}
