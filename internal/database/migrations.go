package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Migrate() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	user := viper.GetString("DB_USER")
	url := viper.GetString("DB_URL")
	port := viper.GetString("DB_PORT")
	database := viper.GetString("DB_DATABASE")
	sslmode := viper.GetString("DB_SSLMODE")

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
	defer db.Close()

	err = db.Ping()
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
