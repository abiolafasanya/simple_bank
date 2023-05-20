package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var dbUrl = "postgresql://abiola:abiola78@localhost:5432/simple_bank?sslmode=disable"
// var ndb = ""postgres://localhost:5432/database?sslmode=disable""

func main() {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		// "./db/migration/000001_init_schema.up.sql",
        "file://db/migration0",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
