package entity

import (
	"time"
)

type CartItem struct {
	tableName        struct{}        `pg:"cart_items,alias:ci"`
	ID               string          `pg:"id,pk"`
	CartID           string          `pg:"cart_id"`
	ProductID        string          `pg:"product_id"`
	ProductVariantID string          `pg:"product_variant_id"`
	Product          *Product        `pg:"rel:has-one"`
	ProductVariant   *ProductVariant `pg:"rel:has-one"`
	Quantity         int             `pg:"quantity"`
	Price            float64         `pg:"price"`
	CreatedAt        time.Time       `pg:"created_at"`
	UpdatedAt        time.Time       `pg:"updated_at"`
}

func (ci *CartItem) NameTable() any {
	return ci.tableName
}
