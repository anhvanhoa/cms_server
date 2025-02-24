package entity

import (
	"time"
)

type MailStatus struct {
	table     struct{}  `pg:"mail_status,alias:ms"`
	Status    string    `pg:"status,pk"`
	Name      string    `pg:"name"`
	CreatedAt time.Time `pg:"created_at"`
}

func (ms *MailStatus) GetNameTable() any {
	return ms.table
}
