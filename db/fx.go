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

func NewHostRepo(db *sql.DB) *HostRepo {
	return &HostRepo{db: db}
}

func NewScanRepo(db *sql.DB) *ScanRepo {
	return &ScanRepo{db: db}
}

func NewDbListener(s ScanRepo, h HostRepo) api.Listener {
	return &DbListener{scans: s, hosts: h}
}

// Functions
func Bootstrap(db *sql.DB) {
	initSchema(db)
}
