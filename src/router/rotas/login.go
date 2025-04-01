package rotas

import (
	"net/http"

	"github.com/nfdeveloper/dev_book_api/src/controllers"
)

var rotaLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
