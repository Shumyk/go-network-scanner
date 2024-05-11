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
	return all(h.db, queryHostSelectAll, hostMap)
}

func (h *HostRepo) Upsert(host api.Host) (int, error) {
	return upsert(
		h.db,
		queryHostUpsert,
		sql.Named("session", h.conf.Name),
		sql.Named("ip", host.IP),
		sql.Named("osType", host.OsType),
		sql.Named("hostnameLabel", host.HostnameLabel),
		sql.Named("customLabel", host.CustomLabel),
	)
}

func hostMap(f scanFunc) (api.Host, error) {
	var h api.Host
	err := f(
		&h.IP,
		&h.FirstDiscovered,
		&h.LastSeen,
		&h.OsType,
		&h.HostnameLabel,
		&h.CustomLabel,
	)
	return h, err
}
