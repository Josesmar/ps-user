package main

import (
	"log"
	"net/http"
	"os"
	"ps-user/src/adapter/api/domain/configuration"
	"ps-user/src/adapter/api/domain/router"
)

func main() {

	configuration.Load()
	r := router.Generate()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe("127.0.0.1:"+port, r); err != nil {
		log.Fatal(err)
	}

}
