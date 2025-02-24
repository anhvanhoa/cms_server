package entity

import (
	"time"
)

type Category struct {
	tableName   struct{}   `pg:"categories,alias:c"`
	ID          string     `pg:"id,pk"`
	Name        string     `pg:"name"`
	Slug        string     `pg:"slug"`
	Description string     `pg:"description"`
	Thumbnail   string     `pg:"thumbnail"`
	ViewCount   int        `pg:"view_count"`
	Type        string     `pg:"type"`
	Status      string     `pg:"status"`
	ParentID    string     `pg:"parent_id"`
	CreatedBy   string     `pg:"created_by"`
	CreatedAt   time.Time  `pg:"created_at"`
	UpdatedAt   *time.Time `pg:"updated_at"`
}

func (c *Category) NameTable() any {
	return c.tableName
}
