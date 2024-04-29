package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func connect() {
	var err error
	db, err = sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	initSchema()
}

func initSchema() {
	exec(create_table_config)
	exec(create_table_scan)
	exec(create_table_host)
}

func exec(query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Panic("Failed to execute: ", query, err)
	}
}
