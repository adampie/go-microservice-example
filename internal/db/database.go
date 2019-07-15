package db

import (
	"database/sql"
	"go.uber.org/zap"
	"os"
)

var db *sql.DB

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

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

	zap.S().Debug("DB_USER: ", user)
	zap.S().Debug("DB_URL: ", url)
	zap.S().Debug("DB_PORT: ", port)
	zap.S().Debug("DB_DATABASE: ", database)
	zap.S().Debug("DB_SSLMODE: ", sslmode)

	dbUrl := "postgres://" + user + "@" + url + ":" + port + "/" + database + "?sslmode=" + sslmode

	var err error

	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		zap.S().Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		zap.S().Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
