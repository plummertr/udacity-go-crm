package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 3000

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	fmt.Printf("Starting server on port %v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
