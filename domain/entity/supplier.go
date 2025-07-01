package entity

import (
	"time"
)

type Supplier struct {
	tableName      struct{}   `pg:"suppliers,alias:s"`
	ID             string     `pg:"id,pk"`
	CreatedBy      string     `pg:"created_by"`
	Name           string     `pg:"name"`
	Type           string     `pg:"type"`
	Service        string     `pg:"service"`
	Address        string     `pg:"address"`
	TaxCode        string     `pg:"tax_code"`
	Representative string     `pg:"representative"`
	Phone          string     `pg:"phone"`
	Email          string     `pg:"email"`
	BankAccount    string     `pg:"bank_account"`
	BankName       string     `pg:"bank_name"`
	BankNumber     string     `pg:"bank_number"`
	Note           string     `pg:"note"`
	Status         string     `pg:"status"`
	CreatedAt      time.Time  `pg:"created_at"`
	UpdatedAt      *time.Time `pg:"updated_at"`
}

func (s *Supplier) NameTable() any {
	return s.tableName
}
