package adapter

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"ps-user/internal/adapter/config"
	"ps-user/internal/adapter/consumer/postgres"
	"ps-user/internal/adapter/producer"
	"ps-user/internal/application/service"
)

func Run() error {

	config.Load()

	db, err := sql.Open("postgres", config.StringConectionBanco)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco de dados:", err)
		return err
	}

	repository := postgres.NewUserRepository(db)
	service := service.NewUserService(repository)
	controller := producer.NewController(service)
	router := producer.NewRouter().RegisterRoutes(controller)

	router = config.Middleware(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
	return nil
}
