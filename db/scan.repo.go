package db

import (
	"database/sql"
	"go-network-scanner/api"
)

type ScanRepo struct {
	db *sql.DB
}

func (s *ScanRepo) All() ([]api.Scan, error) {
	rows, err := s.db.Query(SCAN_SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scans []api.Scan
	for rows.Next() {
		var scan api.Scan
		err := rows.Scan(&scan.Ip, &scan.Port, &scan.FirstDiscovered, &scan.LastSeen, &scan.ServiceBanner)
		if err != nil {
			return nil, err
		}
		scans = append(scans, scan)
	}
	return scans, nil
}

func (s *ScanRepo) Upsert(scan api.Scan) (int, error) {
	result, err := s.db.Exec(SCAN_UPSERT, scan.Ip, scan.Port, scan.ServiceBanner)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
