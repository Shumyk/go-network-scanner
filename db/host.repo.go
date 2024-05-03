package db

import (
	"database/sql"
	"go-network-scanner/api"
)

type HostRepo struct {
	db *sql.DB
}

func (h *HostRepo) All() ([]api.Host, error) {
	rows, err := h.db.Query(HOST_SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hosts []api.Host
	for rows.Next() {
		var h api.Host
		err := rows.Scan(&h.Ip, &h.FirstDiscovered, &h.LastSeen, &h.OsType, &h.HostnameLabel, &h.CustomLabel)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, h)
	}
	return hosts, nil
}

func (h *HostRepo) Insert(host api.Host) (int, error) {
	result, err := h.db.Exec(HOST_INSERT, host.Ip, host.OsType, host.HostnameLabel, host.CustomLabel)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
