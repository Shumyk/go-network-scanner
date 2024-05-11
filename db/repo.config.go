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

// Insert inserts a configuration.
func (c *ConfigRepo) Insert(config *api.Config) (int, error) {
	return upsert(
		c.db,
		queryConfigInsert,
		config.Name, config.EnableProxy,
		config.IPRange1, config.IPRange2, config.IPRange3,
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
