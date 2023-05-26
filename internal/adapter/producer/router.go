package producer

import (
	"github.com/go-chi/cors"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

func (r *Router) RegisterRoutes(controller *Controller) *mux.Router {
	r.router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.router.HandleFunc("/users/{userID}", controller.GetUser).Methods("GET")
	r.router.HandleFunc("/users", controller.GetAllUser).Methods("GET")
	r.router.HandleFunc("/users/credentials/", controller.ValidCredentials).Methods("GET")
	r.router.HandleFunc("/users/{userID}", controller.DeleteUser).Methods("DELETE")
	r.router.HandleFunc("/users", controller.DeleteUsers).Methods("DELETE")
	r.router.HandleFunc("/users/{userID}", controller.UpdateUser).Methods("PATCH")
	return r.router
}

func (r *Router) Generate() *mux.Router {
	rMux := mux.NewRouter()
	rMux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	return r.router
}
