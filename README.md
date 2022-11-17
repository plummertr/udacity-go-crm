# udacity-go-crm
Final project for Udacity Golang course

Server with five api endpoints to manager customer objects. Data is stored in a map (map[string]Customer).  Api endpoints return json.  Home route returns a basic html page listing the endpoints.

## Setup
main.go is in the root of the project.  With any luck, the dependencies will download on first run.  If not, go mod tidy should work.  Only additional dependencies are httprouter and uuid.

## Routes
* Route Listing:     [GET]    /
* Get all customers: [GET]    /customers
* Get one customer:  [GET]    /customers/:id
* Create customer:   [POST]   /customers
* Update customer:   [PUT]    /customers/:id
* Delete customer:   [Delete] /customers/:id
