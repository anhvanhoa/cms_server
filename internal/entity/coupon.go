package entity

import (
	"time"
)

type Coupon struct {
	tableName      struct{}   `pg:"coupons,alias:c"`
	ID             string     `pg:"id,pk"`
	CreatedBy      string     `pg:"created_by"`
	Name           string     `pg:"name"`
	Code           *string    `pg:"code"`
	Type           string     `pg:"type"`
	Value          float64    `pg:"value"`
	Quantity       int        `pg:"quantity"`
	Description    string     `pg:"description"`
	TimeStart      time.Time  `pg:"time_start"`
	TimeEnd        time.Time  `pg:"time_end"`
	MinOrderValue  float64    `pg:"min_order_value"`
	MaxValue       float64    `pg:"max_value"`
	LimitPreUser   int        `pg:"limit_pre_user"`
	Status         string     `pg:"status"`
	CouponForType  string     `pg:"coupon_for_type"`
	RefIds         []string   `pg:"ref_ids,array"`
	PaymentMethods []string   `pg:"payment_methods,array"`
	CreatedAt      time.Time  `pg:"created_at"`
	UpdatedAt      *time.Time `pg:"updated_at"`
}

func (c *Coupon) NameTable() any {
	return c.tableName
}
