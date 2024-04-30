package db

import (
	"database/sql"
	"fmt"
	log "log/slog"

	_ "modernc.org/sqlite"
)

func initSchema(db *sql.DB) {
	exec(db, create_table_config)
	exec(db, create_table_scan)
	exec(db, create_table_host)
}

func exec(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Error(
			"Failed to execute query",
			log.String("query", query),
			log.Any("error", err),
		)
	}
}

func open(databaseName string) *sql.DB {
	var err error
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%v?cache=shared", databaseName))
	if err != nil {
		panic(err)
	}
	return db
}
