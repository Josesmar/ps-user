package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequestAutentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method, "OPTIONS")
	}
	return r
}
