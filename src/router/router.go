package router

import (
	"github.com/gorilla/mux"
	"github.com/nfdeveloper/dev_book_api/src/router/rotas"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
