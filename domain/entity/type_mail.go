package entity

import (
	"time"
)

type TypeMail struct {
	tableName struct{}   `pg:"type_mails,alias:tm"`
	ID        string     `pg:"id,pk"`
	Name      string     `pg:"name,unique"`
	CreatedBy string     `pg:"created_by"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (tm *TypeMail) GetNameTable() any {
	return tm.tableName
}
