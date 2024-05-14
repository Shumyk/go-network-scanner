package db

import (
	"database/sql"
	"go-network-scanner/api"
)

// ScanRepo is a repository for the scan table.
type ScanRepo struct {
	db   *sql.DB
	conf *api.Config
}

// All retrieves all scans.
func (s *ScanRepo) All() ([]*api.Scan, error) {
	return all(s.db, queryScanSelectAll, scanMap)
}

// Upsert inserts or updates a scan.
func (s *ScanRepo) Upsert(scan *api.Scan) (int, error) {
	return upsert(
		s.db,
		queryScanUpsert,
		sql.Named("session", s.conf.Name),
		sql.Named("ip", scan.IP),
		sql.Named("port", scan.Port),
		sql.Named("serviceBanner", scan.ServiceBanner),
	)
}

func scanMap(f scanFunc) (*api.Scan, error) {
	var s api.Scan
	err := f(
		&s.Session,
		&s.IP,
		&s.Port,
		&s.FirstDiscovered,
		&s.LastSeen,
		&s.ServiceBanner,
	)
	return &s, err
}
