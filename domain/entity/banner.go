package entity

import (
	"time"
)

type Banner struct {
	tableName struct{}   `pg:"banners,alias:b"`
	ID        string     `pg:"id,pk"`
	Title     string     `pg:"title"`
	Content   string     `pg:"content"`
	Image     string     `pg:"image"`
	Url       string     `pg:"url"`
	Order     int        `pg:"order"`
	Target    string     `pg:"target"`
	CreatedBy string     `pg:"created_by"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (b *Banner) NameTable() any {
	return b.tableName
}
