package controllers

import (
	"agenda-online/src/autenticacao"
	"agenda-online/src/banco"
	"agenda-online/src/cookies"
	"agenda-online/src/modelos"
	"agenda-online/src/repositorios"
	"agenda-online/src/respostas"
	"agenda-online/src/seguranca"
	"agenda-online/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PaginaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login", nil)
}

func LogarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)
	usuarioNoBanco, erro := repositorio.BuscarUsuarioNoBanco(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = cookies.Salvar(w, token, r); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, nil)
		return
	}

}

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.RepositorioUsuario(db)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	fmt.Println(usuarioId)
	respostas.JSON(w, http.StatusCreated, usuario)
}
