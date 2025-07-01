package entity

import "time"

type ProductVariant struct {
	tableName    struct{}   `pg:"product_variants,alias:pv"`
	ID           string     `pg:"id,pk"`
	ProductID    string     `pg:"product_id"`
	Product      *Product   `pg:"rel:has-one"`
	Sku          string     `pg:"sku"`
	Price        float64    `pg:"price"`
	Discount     float64    `pg:"discount"`
	TypeDiscount string     `pg:"type_discount"`
	Quantity     int        `pg:"quantity"`
	Thumbnail    string     `pg:"thumbnail"`
	Weight       float64    `pg:"weight"`
	Length       float64    `pg:"length"`
	Width        float64    `pg:"width"`
	Height       float64    `pg:"height"`
	CreatedAt    time.Time  `pg:"created_at"`
	UpdatedAt    *time.Time `pg:"updated_at"`
}

func (pv *ProductVariant) NameTable() any {
	return pv.tableName
}
