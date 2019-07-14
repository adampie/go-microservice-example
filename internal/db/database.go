package db

import (
	"database/sql"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sql.DB

func init() {
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
