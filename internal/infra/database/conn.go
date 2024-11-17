package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Abre a conexão
	db, err := sql.Open("mysql", "root:robot@tcp(localhost:3306)/products?parseTime=true")
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão é válida
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
