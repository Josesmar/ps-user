package main

import (
	"log"
	"net/http"
	"os"
	"ps-user/src/adapter/api/domain/configuration"
	"ps-user/src/adapter/api/domain/middleware"
	"ps-user/src/adapter/api/domain/router"
)

func main() {

	configuration.Load()
	router := router.Generate()
	router = middleware.Middleware(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}
