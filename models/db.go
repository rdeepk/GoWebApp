package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=test1 password=tester dbname=goapp sslmode=disable")
	return db, err
}
