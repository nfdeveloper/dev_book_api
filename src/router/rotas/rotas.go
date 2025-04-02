package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nfdeveloper/dev_book_api/src/middlewares"
)

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.Uri, middlewares.Autenticar(rota.Funcao)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
		}
	}

	return r
}
