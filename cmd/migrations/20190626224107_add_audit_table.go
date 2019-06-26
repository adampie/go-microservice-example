package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20190626224107, Down20190626224107)
}

func Up20190626224107(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func Down20190626224107(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
