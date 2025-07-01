package entity

import (
	"time"
)

type Post struct {
	tableName   struct{}   `pg:"posts,alias:p"`
	ID          string     `pg:"id,pk"`
	Title       string     `pg:"title"`
	Slug        string     `pg:"slug"`
	Content     string     `pg:"content"`
	TypeContent string     `pg:"type_content"`
	Thumbnail   string     `pg:"thumbnail"`
	CategoryId  string     `pg:"category_id"`
	Category    *Category  `pg:"rel:has-one"`
	ViewCount   int        `pg:"view_count"`
	Order       int        `pg:"order"`
	Target      string     `pg:"target"`
	Media       []string   `pg:"media,array"`
	Status      string     `pg:"status"`
	CreatedBy   string     `pg:"created_by"`
	Author      *User      `pg:"rel:has-one"`
	CreatedAt   time.Time  `pg:"created_at"`
	UpdatedAt   *time.Time `pg:"updated_at"`
}

func (p *Post) NameTable() any {
	return p.tableName
}
