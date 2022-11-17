package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/plummertr/udacity-go-crm/internal/apiutility"
	"github.com/plummertr/udacity-go-crm/internal/data"
)

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func (a *App) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := a.Models.Customers.GetAll()

	apiutility.WriteJson(w, customers, http.StatusOK)
}

func (a *App) getCustomer(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	customer := a.Models.Customers.Get(id)

	if customer == nil {
		apiutility.WriteJson(w, apiutility.Wrap("error", "Customer not found"), http.StatusNotFound)
		return
	}

	apiutility.WriteJson(w, customer, http.StatusOK)
}

func (a *App) addCustomer(w http.ResponseWriter, r *http.Request) {
	var input data.Customer

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiutility.WriteJson(w, apiutility.Wrap("error", "Unable to read request body"), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		apiutility.WriteJson(w, apiutility.Wrap("error", "invalid data, unable to parse json"), 422)
		return
	}

	id := a.Models.Customers.Create(input)

	if id == nil {
		apiutility.WriteJson(w, apiutility.Wrap("error", "Customer not found"), http.StatusNotFound)
		return
	}

	apiutility.WriteJson(w, a.Models.Customers.GetAll(), http.StatusCreated)
}

func (a *App) updateCustomer(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	var input data.Customer

	json.NewDecoder(r.Body).Decode(&input)
	input.ID = id

	customer := a.Models.Customers.Get(id)

	if customer == nil {
		apiutility.WriteJson(w, apiutility.Wrap("error", "Customer not found"), http.StatusNotFound)
		return
	}

	a.Models.Customers.Update(id, input)

	apiutility.WriteJson(w, a.Models.Customers.GetAll(), http.StatusOK)
}

func (a *App) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	customer := a.Models.Customers.Get(id)

	if customer == nil {
		apiutility.WriteJson(w, apiutility.Wrap("error", "Customer not found"), http.StatusNotFound)
		return
	}

	a.Models.Customers.Delete(id)

	apiutility.WriteJson(w, a.Models.Customers.GetAll(), http.StatusOK)
}
