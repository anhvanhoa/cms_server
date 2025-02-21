package entity

type User struct {
	tableName struct{} `pg:"users,alias:u"`
	ID        string   `pg:"id,pk"`
	Email     string   `pg:"email,unique"`
	Phone     string   `pg:"phone,unique"`
	Password  string   `pg:"password"`
	FullName  string   `pg:"full_name"`
}

type UserInfor struct {
	ID       string `pg:"id,pk"`
	Email    string `pg:"email,unique"`
	Phone    string `pg:"phone,unique"`
	FullName string `pg:"full_name"`
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetNameTable() any {
	return u.tableName
}

func (u *User) GetInfor() UserInfor {
	return UserInfor{
		ID:       u.ID,
		Email:    u.Email,
		Phone:    u.Phone,
		FullName: u.FullName,
	}
}
