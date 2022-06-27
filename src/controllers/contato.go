package controllers

import (
	"agenda-online/src/autenticacao"
	"agenda-online/src/banco"
	"agenda-online/src/modelos"
	"agenda-online/src/repositorios"
	"agenda-online/src/respostas"
	"agenda-online/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CadastrarContato(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var contato modelos.Contato
	if erro = json.Unmarshal(corpoRequisicao, &contato); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		return
	}

	repositorio := repositorios.RepositorioContato(db)
	repositorio.CriarContato(contato, usuarioID)

	http.Redirect(w, r, "/home", 301)
}

func Home(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro conectar"))
		return
	}
	defer db.Close()

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		return
	}

	linhas, erro := db.Query("select * from contatos where usuario_dono = ?", usuarioID)
	if erro != nil {
		w.Write([]byte("erro busca"))
		return
	}
	defer linhas.Close()

	contato := modelos.Contato{}
	arrayContato := []modelos.Contato{}

	for linhas.Next() {
		var id, dono uint64
		var nome, apelido, site, email, telefone, endereco string

		if erro := linhas.Scan(&id, &dono, &nome, &apelido, &site, &email, &telefone, &endereco); erro != nil {
			return
		}

		contato.ID = id
		contato.Dono = dono
		contato.Nome = nome
		contato.Apelido = apelido
		contato.Site = site
		contato.Email = email
		contato.Telefone = telefone
		contato.Endereco = endereco

		arrayContato = append(arrayContato, contato)

	}

	utils.ExecutarTemplate(w, "home", arrayContato)

}

func CadastrarContatoHtml(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastrar", nil)
}

func BuscarContato(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		fmt.Println("erro id")
		return
	}

	parametros := mux.Vars(r)

	contatoID, erro := strconv.ParseUint(parametros["contatoID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.RepositorioContato(db)
	contato, erro := repositorio.BuscarPorID(contatoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if contato.Dono != usuarioID {
		http.Redirect(w, r, "/home", 302)
		return

	}

	utils.ExecutarTemplate(w, "editar", contato)
}

func AtualizarContato(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	contatoID, erro := strconv.ParseUint(parametros["contatoID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.RepositorioContato(db)
	contatoNoBanco, erro := repositorio.BuscarPorID(contatoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioIDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		fmt.Println("erro id")
		return
	}

	fmt.Println(usuarioIDToken)
	fmt.Println(contatoNoBanco.Dono)

	if contatoNoBanco.Dono != usuarioIDToken {
		http.Redirect(w, r, "/home", 302)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var contato modelos.Contato
	if erro = json.Unmarshal(corpoRequisicao, &contato); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorio.Atualizar(contatoID, contato); erro != nil {
		return
	}

	http.Redirect(w, r, "/home", 302)

}

func DeletarContato(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		fmt.Println("erro id")
		return
	}

	contatoID, erro := strconv.ParseUint(parametros["contatoID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.RepositorioContato(db)
	contato, erro := repositorio.BuscarPorID(contatoID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if contato.Dono != usuarioID {
		http.Redirect(w, r, "/home", 302)
		return

	}

	erro = repositorio.DeletarContato(contatoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	http.Redirect(w, r, "/home", 302)
}
