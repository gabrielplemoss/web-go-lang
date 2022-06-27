package banco

import (
	"agenda-online/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringBanco)

	if erro != nil {
		return nil, erro
	}
	
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	

	return db, nil
}
