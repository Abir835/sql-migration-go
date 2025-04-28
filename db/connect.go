package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	// Change connection string based on your database
	db, err = sql.Open("postgres", "postgres://username:password@localhost:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	return db
}
