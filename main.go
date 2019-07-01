package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"go.uber.org/zap"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func init() {
	if os.Getenv("ENV") == "PRODUCTION" {
		logger, _ := zap.NewProduction()
		zap.ReplaceGlobals(logger)
	} else {
		logger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
	}
}

func main() {
	user := os.Getenv("DB_USER")
	url := os.Getenv("DB_URL")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	sslmode := os.Getenv("DB_SSLMODE")

	if user == "" {
		user = "postgres"
	}
	if url == "" {
		url = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if database == "" {
		database = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	dbUrl := "postgres://" + user + "@" + url + ":" + port + "/" + database + "?sslmode=" + sslmode

	db, err := sql.Open("postgres", dbUrl)
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
