package db

import (
	"database/sql"
	"go-network-scanner/api"
)

// Constructors
func NewDb(databaseName string) any {
	return func() *sql.DB { return open(databaseName) }
}

func NewConfigRepo(db *sql.DB, conf *api.Config) *ConfigRepo {
	repo := &ConfigRepo{db: db}
	repo.Insert(*conf)
	repo.Current = conf
	return repo
}

// Functions
func Bootstrap(db *sql.DB) {
	initSchema(db)
}
