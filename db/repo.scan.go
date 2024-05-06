package db

import (
	"database/sql"
	"go-network-scanner/api"
)

type ScanRepo struct {
	db *sql.DB
}

func (s *ScanRepo) All() ([]api.Scan, error) {
	return all(s.db, SCAN_SELECT_ALL, scanMap)
}

func (s *ScanRepo) Upsert(scan api.Scan) (int, error) {
	return upsert(
		s.db,
		SCAN_UPSERT,
		sql.Named("ip", scan.Ip),
		sql.Named("port", scan.Port),
		sql.Named("serviceBanner", scan.ServiceBanner),
	)
}

func scanMap(f func(...any) error) (api.Scan, error) {
	var s api.Scan
	err := f(
		&s.Ip,
		&s.Port,
		&s.FirstDiscovered,
		&s.LastSeen,
		&s.ServiceBanner,
	)
	return s, err
}
