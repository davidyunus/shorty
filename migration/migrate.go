package main

import (
	"database/sql"
	"log"

	"github.com/davidyunus/shorty/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", config.DBConnectionString())
	if err != nil {
		log.Fatal("error when open postgres connection: ", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("error when creating postgres instance: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migration", "postgres", driver)

	if err != nil {
		log.Fatal("error when creating database instance: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("error when migrate up: ", err)
	}
}
