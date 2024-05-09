package db

import (
	"database/sql"
	"go-network-scanner/api"
)

// Constructors.
func NewDb(databaseName string) any {
	return func() *sql.DB { return open(databaseName) }
}

func NewConfigRepo(db *sql.DB, conf *api.Config) *ConfigRepo {
	repo := &ConfigRepo{db, conf}
	return repo
}

func NewHostRepo(db *sql.DB, conf *api.Config) *HostRepo {
	return &HostRepo{db, conf}
}

func NewScanRepo(db *sql.DB, conf *api.Config) *ScanRepo {
	return &ScanRepo{db, conf}
}

func NewDbHostListener(s *ScanRepo) api.Listener {
	return &ScanListener{scans: s}
}

func NewDbScanListener(h *HostRepo) api.Listener {
	return &HostListener{hosts: h}
}

// Functions.
func Bootstrap(db *sql.DB, confRepo *ConfigRepo) {
	initSchema(db)
	confRepo.Insert(*confRepo.Current)
}
