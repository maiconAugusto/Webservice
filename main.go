package main

import (
	"app/src/routes"
	"log"
	"net/http"
)

type name struct {
	Name string `json:"name"`
}

func main() {
	port := ":8080"
	log.Fatal(http.ListenAndServe(port, routes.Routes()))
}
