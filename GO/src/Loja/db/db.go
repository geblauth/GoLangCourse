package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBanco() *sql.DB {
	conexao := "user=postgres dbname=germano_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic((err.Error()))
	}
	return db

}