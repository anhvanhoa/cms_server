package entity

import "time"

type StatusHistory struct {
	tableName     struct{}   `pg:"status_histories,alias:sh"`
	Status        StatusMail `pg:"status"`
	MailHistoryId string     `pg:"mail_history_id"`
	Message       string     `pg:"message"`
	CreatedAt     time.Time  `pg:"created_at"`
}

func (sh *StatusHistory) GetNameTable() any {
	return sh.tableName
}
