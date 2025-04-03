package repositorios

import (
	"database/sql"

	"github.com/nfdeveloper/dev_book_api/src/modelos"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick 
		from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?
	`, publicacaoId)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutoNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p
	inner join usuarios u on u.id = p.autor_id
	inner join seguidores s on p.autor_id = s.usuario_id
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc`,
		usuarioId, usuarioId,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	if linhas.Next() {

		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutoNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}
