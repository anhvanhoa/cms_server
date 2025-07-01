package entity

import (
	"time"
)

type StatusMail string

const (
	MAIL_STATUS_PENDING   StatusMail = "pending"
	MAIL_STATUS_SENT      StatusMail = "sent"
	MAIL_STATUS_DELIVERED StatusMail = "delivered"
	MAIL_STATUS_FAILED    StatusMail = "failed"
	MAIL_STATUS_CANCELED  StatusMail = "canceled"
	MAIL_STATUS_CLICKED   StatusMail = "clicked"
	MAIL_STATUS_OPENED    StatusMail = "opened"
)

type MailStatus struct {
	table     struct{}   `pg:"mail_status,alias:ms"`
	Status    StatusMail `pg:"status,pk"`
	Name      string     `pg:"name"`
	CreatedAt time.Time  `pg:"created_at"`
}

func (ms *MailStatus) GetNameTable() any {
	return ms.table
}
