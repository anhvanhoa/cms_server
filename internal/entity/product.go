package entity

import (
	"time"
)

type Product struct {
	tableName      struct{}   `pg:"products,alias:p"`
	ID             string     `pg:"id,pk"`
	CreatedBy      string     `pg:"created_by"`
	Name           string     `pg:"name"`
	Description    string     `pg:"description"`
	CategoryID     string     `pg:"category_id"`
	Category       *Category  `pg:"rel:has-one"`
	Thumbnail      string     `pg:"thumbnail"`
	Images         []string   `pg:"images,array"`
	Width          float32    `pg:"width"`
	Height         float32    `pg:"height"`
	Weight         float32    `pg:"weight"`
	Length         float32    `pg:"length"`
	Condition      string     `pg:"condition"` // new, used
	Type           string     `pg:"type"`      // physical, digital
	PerOrder       bool       `pg:"per_order"`
	Sku            string     `pg:"sku"`
	MethodPayments []string   `pg:"method_payment,array"`
	Status         string     `pg:"status"`
	CreatedAt      time.Time  `pg:"created_at"`
	UpdatedAt      *time.Time `pg:"updated_at"`
}

func (p *Product) NameTable() any {
	return p.tableName
}
