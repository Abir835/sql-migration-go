package db

import (
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
)

func MigrationDB() {

	db := connectDB()

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Applied %d migrations!\n", n)
}
