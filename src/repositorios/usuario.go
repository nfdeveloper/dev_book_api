package repositorios

import (
	"database/sql"

	"github.com/nfdeveloper/dev_book_api/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
