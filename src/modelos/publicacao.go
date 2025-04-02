package modelos

import "time"

type Publicacao struct {
	ID       uint64    `json:"id, omitempty"`
	Titulo   string    `json:"titulo"`
	Conteudo string    `json:"conteudo"`
	AutorId  uint64    `json:"autoId"`
	AutoNick string    `json:"autorNick"`
	Curtidas uint64    `json:"curtidas"`
	CriadaEm time.Time `json:"criadaEm,omitempty"`
}
