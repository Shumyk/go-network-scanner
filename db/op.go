package db

import (
	"database/sql"
	"fmt"
	log "log/slog"

	// Import the sqlite driver.
	_ "modernc.org/sqlite"
)

//nolint:errcheck,revive // ignore errors because they are handled in the exec()
func initSchema(db *sql.DB) {
	exec(db, CREATE_TABLE_CONFIG)
	exec(db, CREATE_TABLE_SCAN)
	exec(db, CREATE_TABLE_HOST)
}

func exec(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		log.Error(
			"failed to execute query",
			log.String("query", query),
			log.Any("error", err),
		)
		return fmt.Errorf(
			"failed to execute query [%v]: %w",
			query, err,
		)
	}
	return nil
}

func open(databaseName string) *sql.DB {
	databaseName = fmt.Sprintf("file:%v", databaseName)
	db, err := sql.Open("sqlite", databaseName)
	if err != nil {
		panic(err)
	}
	return db
}
