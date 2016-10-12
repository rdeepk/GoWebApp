package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=goapp sslmode=verify-full")
	return db, err
}
