package main

import (
	"flag"

	"github.com/plummertr/udacity-go-crm/api"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 3000, "Provide port for server.  Default is 3000")
	flag.Parse()

	app := api.NewApp(port)

	app.Serve()
}
