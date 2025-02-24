package entity

import "time"

type Menu struct {
	tableName struct{}   `pg:"menus,alias:m"`
	ID        string     `pg:"id,pk"`
	Name      string     `pg:"name"`
	Url       string     `pg:"url"`
	Icon      string     `pg:"icon"`
	ParentID  string     `pg:"parent_id"`
	Order     int        `pg:"order"`
	Target    string     `pg:"target"`
	Image     string     `pg:"image"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (m *Menu) NameTable() any {
	return m.tableName
}
