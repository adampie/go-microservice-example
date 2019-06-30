package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
)

func main() {

	if os.Getenv("ENV") == "PRODUCTION" {
		logger, _ := zap.NewProduction()
		zap.ReplaceGlobals(logger)
	} else {
		logger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
	}

	db, err := sql.Open("postgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		zap.S().Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		zap.S().Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		zap.S().Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		zap.S().Fatal(err)
	}
}


