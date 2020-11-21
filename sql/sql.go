package sql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func OpenDbConnection() {
	db, err = sql.Open("postgres", "user=postgres host=localhost dbname=tarjeta sslmode=disable")
	
	if err != nil {
		log.Fatal(err)
	}
}

func CrearDB() {
	_db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer _db.Close()

	_, err = _db.Exec(`CREATE DATABASE tarjeta`)
	if err != nil {
		log.Fatal(err)
	}
	
}

