package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}
