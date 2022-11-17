# udacity-go-crm
Final project for Udacity Go course

Server with five api endpoints to manager customer objects. Data is stored in a map (map[string]Customer).  Api endpoints return json.  Home route returns a basic html page listing the endpoints.

## Setup
main.go is in the root of the project.  With any luck, the dependencies will download on first run.  If not, go mod tidy should work.  Only additional dependencies are httprouter and uuid.

After cloning the project, go run main.go should be all that's required to run the api.

## Routes
All routes return an array of customers with the exception of / and GET /customers/:id. An error object will be returned in the case of an invalid json body or an id isn't found in the map.
* Route Listing:     [GET]    /
* Get all customers: [GET]    /customers
* Get one customer:  [GET]    /customers/:id
* Create customer:   [POST]   /customers
* Update customer:   [PUT]    /customers/:id
* Delete customer:   [Delete] /customers/:id
