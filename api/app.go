package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/plummertr/udacity-go-crm/internal/data"
)

type App struct {
	Port   int
	Models data.Models
}

func NewApp(port int) *App {
	a := &App{Port: port}
	a.Models = data.NewModels(a.SeedData())

	return a
}

func (a *App) Serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", a.Port),
		Handler: a.Routes(),
	}

	fmt.Printf("Starting the server on port %v\n", a.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
