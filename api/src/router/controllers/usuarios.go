package controllers

import "net/http"

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

// BuscarUsuarios busta todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

// BuscarUsuario busca um úsuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar único usuário!"))
}

// AtualizandoUsuario altera as informações de um úsuario no banco
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}

// DeletarUsuario excluí as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
