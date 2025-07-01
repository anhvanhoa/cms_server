package entity

import "time"

type MailHistory struct {
	tableName     struct{}       `pg:"mail_histories,alias:mh"`
	ID            string         `pg:"id,pk"`
	TemplateId    string         `pg:"template_id"`
	Subject       string         `pg:"subject"`
	Body          string         `pg:"body"`
	To            string         `pg:"to"`
	Tos           []string       `pg:"tos"`
	Data          map[string]any `pg:"data"`
	EmailProvider string         `pg:"email_provider"`
	CreatedBy     string         `pg:"created_by"`
	CreatedAt     time.Time      `pg:"created_at"`
	UpdatedAt     *time.Time     `pg:"updated_at"`
}

func (mh *MailHistory) NameTable() any {
	return mh.tableName
}
