package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *App) Routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/customers", a.getCustomers)
	r.HandlerFunc(http.MethodGet, "/customers/:id", a.getCustomer)
	r.HandlerFunc(http.MethodPost, "/customers", a.addCustomer)
	r.HandlerFunc(http.MethodPut, "/customers/:id", a.updateCustomer)
	r.HandlerFunc(http.MethodDelete, "/customers/:id", a.deleteCustomer)

	return r
}
