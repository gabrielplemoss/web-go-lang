package repositorios

import (
	"agenda-online/src/modelos"
	"database/sql"
	"fmt"
)

type Contato struct {
	db *sql.DB
}

func RepositorioContato(db *sql.DB) *Contato {
	return &Contato{db}
}

func (repositorio Contato) CriarContato(contato modelos.Contato, usuarioDonoID uint64) {
	statement, erro := repositorio.db.Prepare(
		"insert into contatos (usuario_dono, nome, apelido, site, email, telefone, endereco) values (?, ?, ?, ?, ?, ?, ?)",
	)

	if erro != nil {
		return
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuarioDonoID, contato.Nome, contato.Apelido, contato.Site, contato.Email, contato.Telefone, contato.Endereco)
	if erro != nil {
		return
	}

	fmt.Println(resultado)

}

func (repositorio Contato) BuscarPorID(contatoID uint64) (modelos.Contato, error) {
	linhas, erro := repositorio.db.Query("select * from contatos where id = ?", contatoID)
	if erro != nil {
		return modelos.Contato{}, erro
	}
	defer linhas.Close()

	var contato modelos.Contato

	if linhas.Next() {
		if erro := linhas.Scan(
			&contato.ID,
			&contato.Dono,
			&contato.Nome,
			&contato.Apelido,
			&contato.Site,
			&contato.Email,
			&contato.Telefone,
			&contato.Endereco,
		); erro != nil {
			return modelos.Contato{}, erro
		}
	}

	return contato, erro
}

func (repositorio Contato) DeletarContato(contatoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from contatos where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(contatoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Contato) Atualizar(contatoID uint64, contato modelos.Contato) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE contatos SET nome = ?, apelido = ?, site = ?, email = ?, telefone = ?, endereco = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(contato.Nome, contato.Apelido, contato.Site, contato.Email, contato.Telefone, contato.Endereco, contatoID); erro != nil {
		return erro
	}

	return nil
}
