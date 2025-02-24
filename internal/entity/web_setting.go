package entity

import (
	"time"
)

type WebSetting struct {
	tableName struct{}   `pg:"web_settings,alias:ws"`
	Key       string     `pg:"key,pk"`
	Value     string     `pg:"value"`
	CreatedBy string     `pg:"created_by"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (ws *WebSetting) NameTable() any {
	return ws.tableName
}
