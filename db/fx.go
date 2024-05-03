package db

import (
	"database/sql"
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
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

// Functions
func Bootstrap(db *sql.DB) {
	initSchema(db)
	// TODO: implement persistence
	events.On("scan", func(data ...interface{}) {
		scan := data[0].(api.ScanFound)
		log.Info("Scan received", log.Any("scan", scan))
	})
}
