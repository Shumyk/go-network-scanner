package db

import (
	"database/sql"
	"fmt"
	"go-network-scanner/api"
	"log"

	_ "modernc.org/sqlite"
)

var (
	db *sql.DB

	Configs *ConfigRepo
)

func Init(databaseName string, config *api.Config) {
	db = open(databaseName)
	initSchema()

	Configs = NewConfigRepo(*config)
}

func open(databaseName string) *sql.DB {
	var err error
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%v?cache=shared", databaseName))
	if err != nil {
		panic(err)
	}
	return db
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
