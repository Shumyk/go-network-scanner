package db

import (
	"database/sql"
	"go-network-scanner/api"
)

type HostRepo struct {
	db   *sql.DB
	conf *api.Config
}

func (h *HostRepo) All() ([]api.Host, error) {
	return all(h.db, HOST_SELECT_ALL, hostMap)
}

func (h *HostRepo) Upsert(host api.Host) (int, error) {
	return upsert(
		h.db,
		HOST_UPSERT,
		sql.Named("session", h.conf.Name),
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
