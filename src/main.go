package main

import (
	"log"
	"net/http"

	"github.com/NullJupiter/GoPresentation/src/templating"

	"github.com/NullJupiter/GoPresentation/src/routing"
)

func main() {
	// Initialize templates
	err := templating.InitTemplating("views/templates/*.html")
	if err != nil {
		log.Fatalf("could not initialize templates: %v", err)
	}

	// Initialize basic routing
	router := routing.InitRouting()

	log.Fatal(http.ListenAndServe(":8080", router))
}
