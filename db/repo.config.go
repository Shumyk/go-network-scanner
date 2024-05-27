package db

import (
	"database/sql"
	"go-network-scanner/api"
)

// ConfigRepo is a repository for the configuration table.
type ConfigRepo struct {
	db      *sql.DB
	Current *api.Config
}

// Get retrieves a configuration by name.
func (c *ConfigRepo) Get(name string) (api.Config, error) {
	return get(c.db, queryConfigSelectByName, mapConfig, name)
}

// All retrieves all configurations.
func (c *ConfigRepo) All() ([]api.Config, error) {
	return all(
		c.db,
		queryConfigSelectAll,
		mapConfig,
	)
}

// Upsert inserts a configuration.
func (c *ConfigRepo) Upsert(config *api.Config) (int, error) {
	return upsert(
		c.db,
		queryConfigUpsert,
		sql.Named("name", config.Name),
		sql.Named("ipRange1", config.IPRange1),
		sql.Named("ipRange2", config.IPRange2),
		sql.Named("ipRange3", config.IPRange3),
		sql.Named("enableProxy", config.EnableProxy),
		sql.Named("proxy1", config.Proxy1),
		sql.Named("proxy2", config.Proxy2),
		sql.Named("proxy3", config.Proxy3),
	)
}

func mapConfig(f scanFunc) (api.Config, error) {
	var c api.Config
	err := f(
		&c.Name,
		&c.EnableProxy,
		&c.IPRange1,
		&c.IPRange2,
		&c.IPRange3,
	)
	return c, err
}
