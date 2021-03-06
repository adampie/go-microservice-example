package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Files read from /migrations
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// Migrate the database
func Migrate() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	db := GetDB()

	defer func() {
		err := db.Close()
		if err != nil {
			zap.S().Fatal(err)
		}
	}()

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
