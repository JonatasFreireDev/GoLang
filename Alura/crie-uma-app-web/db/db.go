package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaComBD() *sql.DB {
	conexao := "user=postgres dbname=loja password=123456 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
