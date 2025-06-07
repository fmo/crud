package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	*sql.DB
}

func NewDatabase(connStr string) *Database {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
	return &Database{db}
}
