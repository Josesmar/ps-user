package router

import (
	"ps-user/internal/users/domain/router/routes"

	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	return routes.Config(r)
}
