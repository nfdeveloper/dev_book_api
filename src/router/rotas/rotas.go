package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
