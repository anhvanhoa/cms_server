package entity

import (
	"time"
)

type OrderStatusHistory struct {
	tableName struct{}   `pg:"order_status_histories,alias:osh"`
	ID        string     `pg:"id,pk"`
	OrderID   string     `pg:"order_id"`
	StatusID  string     `pg:"status_id"`
	Note      string     `pg:"note"`
	CreatedBy string     `pg:"created_by"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (osh *OrderStatusHistory) NameTable() any {
	return osh.tableName
}
