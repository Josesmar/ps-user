package main

import (
	"fmt"
	"log"
	"net/http"
	"ps-user/configuration"
	"ps-user/router"
)

func main() {

	configuration.Load()
	r := router.Generate()

	fmt.Printf("Escutando em porta %s:%d", configuration.IP, configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.IP, configuration.Port), r))
}
