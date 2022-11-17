package data

type Models struct {
	Customers CustomerModel
}

func NewModels(db *map[string]Customer) Models {
	return Models{
		Customers: CustomerModel{DB: db},
	}
}
