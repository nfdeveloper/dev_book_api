package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID       uint64    `json:"id, omitempty"`
	Titulo   string    `json:"titulo"`
	Conteudo string    `json:"conteudo"`
	AutorId  uint64    `json:"autoId"`
	AutoNick string    `json:"autorNick"`
	Curtidas uint64    `json:"curtidas"`
	CriadaEm time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco.")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco.")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
