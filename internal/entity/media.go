package entity

import (
	"time"
)

type Media struct {
	tableName struct{}   `pg:"medias,alias:m"`
	ID        string     `pg:"id,pk"`
	CreatedBy string     `pg:"created_by"`
	Name      string     `pg:"name"`
	MimeType  string     `pg:"mime_type"`
	Size      int64      `pg:"size"`
	Url       string     `pg:"url"`
	Width     float32    `pg:"width"`
	Height    float32    `pg:"height"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (m *Media) NameTable() any {
	return m.tableName
}
