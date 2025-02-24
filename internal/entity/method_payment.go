package entity

type MethodPayment struct {
	tableName struct{} `pg:"method_payments,alias:mp"`
	Id        string   `pg:"id,pk"`
	Name      string   `pg:"name"`
	CreatedAt string   `pg:"created_at"`
	UpdatedAt string   `pg:"updated_at"`
}

func (mp *MethodPayment) NameTable() any {
	return mp.tableName
}
