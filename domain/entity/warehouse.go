package entity

import (
	"time"
)

type Warehouse struct {
	tableName        struct{}   `pg:"warehouses,alias:w"`
	Id               string     `pg:"id,pk"`
	CreatedBy        string     `pg:"created_by"`
	Quantity         int        `pg:"quantity"`
	Total            int        `pg:"total"`
	Name             string     `pg:"name"`
	Cateogty         string     `pg:"category"`
	Unit             string     `pg:"unit"`
	Location         string     `pg:"location"`
	LocatioWarehouse string     `pg:"location_warehouse"`
	Note             string     `pg:"note"`
	Type             string     `pg:"type"`
	StoreName        string     `pg:"store_name"`
	ReceiverName     string     `pg:"receiver_name"`
	ReceiverPhone    string     `pg:"receiver_phone"`
	ReceiverAddress  string     `pg:"receiver_address"`
	SupplierId       string     `pg:"supplier_id"`
	Supplier         *Supplier  `pg:"rel:has-one"`
	CreatedAt        time.Time  `pg:"created_at"`
	UpdatedAt        *time.Time `pg:"updated_at"`
}

func (w *Warehouse) NameTable() any {
	return w.tableName
}
