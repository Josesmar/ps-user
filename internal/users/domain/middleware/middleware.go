package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Middleware(router *mux.Router) *mux.Router {
	cors := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:         []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:         []string{"Link"},
		AllowCredentials:       true,
		OptionsPassthrough:     true,
		MaxAge:                 3599, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)
	return router
}
