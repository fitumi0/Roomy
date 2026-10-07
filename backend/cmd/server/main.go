package main

import (
	"log"
	"net/http"
	"roomy/internal/api"
)

func main() {
	/*
		0) settings parsing
		1) ws hub
		2) goroutine for hub
		3) router
		4) http server
		5) start server
	*/

	r := api.NewRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
