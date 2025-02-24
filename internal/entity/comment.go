package entity

import (
	"time"
)

type Comment struct {
	tableName struct{}  `pg:"comments,alias:c"`
	ID        string    `pg:"id,pk"`
	UserID    string    `pg:"user_id"`
	User      *User     `pg:"rel:has-one"`
	PostID    string    `pg:"post_id"`
	Post      *Post     `pg:"rel:has-one"`
	ParentID  string    `pg:"parent_id"`
	Parent    *Comment  `pg:"rel:has-one"`
	Content   string    `pg:"content"`
	Meida     string    `pg:"media"`
	Status    string    `pg:"status"`
	TypeMedia string    `pg:"type_media"`
	CreatedAt time.Time `pg:"created_at"`
}

func (c *Comment) NameTable() any {
	return c.tableName
}
