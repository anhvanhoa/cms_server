package entity

import (
	"time"
)

type OrderItem struct {
	tableName        struct{}   `pg:"order_items,alias:oi"`
	Id               string     `pg:"id,pk"`
	OrderId          string     `pg:"order_id"`
	Order            *Order     `pg:"rel:has-one"`
	ProductId        string     `pg:"product_id"`
	ProductVariantId string     `pg:"product_variant_id"`
	Name             string     `pg:"name"`
	TypeVariant      string     `pg:"type_variant"`
	Quantity         int        `pg:"quantity"`
	Price            int        `pg:"price"`
	Discount         int        `pg:"discount"`
	TypeDiscount     string     `pg:"type_discount"`
	CreatedAt        time.Time  `pg:"created_at"`
	UpdatedAt        *time.Time `pg:"updated_at"`
}

func (oi *OrderItem) NameTable() any {
	return oi.tableName
}
