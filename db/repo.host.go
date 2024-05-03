package db

import (
	"database/sql"
	"go-network-scanner/api"
)

type HostRepo struct {
	db *sql.DB
}

func (h *HostRepo) All() ([]api.Host, error) {
	return All(h.db, HOST_SELECT_ALL, hostMap)
}

func (h *HostRepo) Upsert(host api.Host) (int, error) {
	return Upsert(
		h.db,
		HOST_UPSERT,
		sql.Named("ip", host.Ip),
		sql.Named("osType", host.OsType),
		sql.Named("hostnameLabel", host.HostnameLabel),
		sql.Named("customLabel", host.CustomLabel),
	)
}

func hostMap(f func(...any) error) (api.Host, error) {
	var h api.Host
	err := f(
		&h.Ip,
		&h.FirstDiscovered,
		&h.LastSeen,
		&h.OsType,
		&h.HostnameLabel,
		&h.CustomLabel,
	)
	return h, err
}
