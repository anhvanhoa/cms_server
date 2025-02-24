package entity

import (
	"time"
)

type PostTag struct {
	tableName struct{}  `pg:"post_tags,alias:pt"`
	PostID    string    `pg:"post_id"`
	TagID     string    `pg:"tag_id"`
	CreatedBy string    `pg:"created_by"`
	CreatedAt time.Time `pg:"created_at"`
}

func (pt *PostTag) NameTable() any {
	return pt.tableName
}
