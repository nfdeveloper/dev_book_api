package middlewares

import (
	"net/http"

	"github.com/nfdeveloper/dev_book_api/src/autenticacao"
	"github.com/nfdeveloper/dev_book_api/src/respostas"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
