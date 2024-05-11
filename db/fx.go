package db

import (
	"database/sql"
	"go-network-scanner/api"
	log "log/slog"
)

// NewDb returns a new database connection.
func NewDb(databaseName string) any {
	return func() *sql.DB { return open(databaseName) }
}

// NewConfigRepo returns a new configuration repository.
func NewConfigRepo(db *sql.DB, conf *api.Config) *ConfigRepo {
	repo := &ConfigRepo{db, conf}
	return repo
}

// NewHostRepo returns a new host repository.
func NewHostRepo(db *sql.DB, conf *api.Config) *HostRepo {
	return &HostRepo{db, conf}
}

// NewScanRepo returns a new scan repository.
func NewScanRepo(db *sql.DB, conf *api.Config) *ScanRepo {
	return &ScanRepo{db, conf}
}

// NewDbHostListener returns a new host event listener.
func NewDbHostListener(h RepoBasic[*api.Host]) api.Listener {
	return &HostListener{hosts: h}
}

// NewDbScanListener returns a new scan event listener.
func NewDbScanListener(s *ScanRepo) api.Listener {
	return &ScanListener{scans: s}
}

// Bootstrap initializes the database.
func Bootstrap(db *sql.DB, confRepo *ConfigRepo) {
	initSchema(db)

	_, err := confRepo.Insert(confRepo.Current)
	if err != nil {
		log.Warn(
			"Failed to insert configuration",
			log.Any("error", err),
		)
	}
}
