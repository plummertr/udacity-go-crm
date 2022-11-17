package data

import "github.com/google/uuid"

type Customer struct {
	ID        string `json:id`
	Name      string `json:name`
	Role      string `json:role`
	Email     string `json:email`
	Phone     string `json:phone`
	Contacted bool   `json:contacted`
}

type CustomerModel struct {
	DB *map[string]Customer
}

func (c CustomerModel) GetAll() map[string]Customer {
	return *c.DB
}

func (c CustomerModel) Get(id string) *Customer {
	if cust, ok := (*c.DB)[id]; ok {

		return &cust
	}

	return nil
}

func (c CustomerModel) Create(cust Customer) *string {
	id := uuid.NewString()

	cust.ID = id

	(*c.DB)[cust.ID] = cust

	return &id
}

func (c CustomerModel) Update(id string, cust Customer) int {
	(*c.DB)[id] = cust

	return 1
}

func (c CustomerModel) Delete(id string) bool {
	delete((*c.DB), id)

	return true
}
