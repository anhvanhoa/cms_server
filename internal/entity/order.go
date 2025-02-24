package entity

import (
	"time"
)

type Order struct {
	tableName         struct{}   `pg:"orders,alias:o"`
	ID                string     `pg:"id,pk"`
	UserID            string     `pg:"user_id"`
	User              *User      `pg:"rel:has-one"`
	Code              string     `pg:"code"`
	FullName          string     `pg:"full_name"`
	Email             string     `pg:"email"`
	Phone             string     `pg:"phone"`
	Address           string     `pg:"address"`
	Location          string     `pg:"location"`
	MethodPayment     string     `pg:"method_payment"`
	PaymentStatus     string     `pg:"payment_status"`
	SubTotal          float64    `pg:"sub_total"`
	Shipping          float64    `pg:"shipping"`
	ShippingDiscount  float64    `pg:"shipping_discount"`
	Note              string     `pg:"note"`
	TotalPriceProduct float64    `pg:"total_price_product"`
	DiscountProduct   float64    `pg:"discount_product"`
	ToAddress         string     `pg:"to_address"`
	ToLocation        string     `pg:"to_location"`
	Total             float64    `pg:"total"`
	CreatedAt         time.Time  `pg:"created_at"`
	UpdatedAt         *time.Time `pg:"updated_at"`
}

func (o *Order) NameTable() any {
	return o.tableName
}
