package router

import (
	"api/src/rotas"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
