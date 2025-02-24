package entity

import (
	"time"
)

type MailProvider struct {
	tableName  struct{}  `pg:"mail_providers,alias:mp"`
	Email      string    `pg:"email,pk"`
	Password   string    `pg:"password"`
	UserName   string    `pg:"user_name"`
	Port       string    `pg:"port"`
	Host       string    `pg:"host"`
	Encryption string    `pg:"encryption"`
	Mailer     string    `pg:"mailer"`
	Name       string    `pg:"name"`
	TypeId     string    `pg:"type_id"`
	Type       TypeMail  `pg:"rel:has-one"`
	CreatedBy  string    `pg:"created_by"`
	CreatedAt  time.Time `pg:"created_at"`
	UpdatedAt  time.Time `pg:"updated_at"`
}

func (mp *MailProvider) GetNameTable() any {
	return mp.tableName
}
