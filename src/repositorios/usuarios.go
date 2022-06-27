package repositorios

import (
	"agenda-online/src/modelos"
	"agenda-online/src/seguranca"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func RepositorioUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}

}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (usuario, email, senha) values (?, ?, ?)",
	)

	if erro != nil {
		return 0, nil
	}

	//defer statement.Close()

	senhaHash, erro := seguranca.Hash(usuario.Senha)
	if erro != nil {
		return 0, nil
	}

	usuario.Senha = string(senhaHash)

	resultado, erro := statement.Exec(usuario.Usuario, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, nil
	}

	ultimoId, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoId), nil
}

func (repositorio usuarios) BuscarUsuarioNoBanco(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	//defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
