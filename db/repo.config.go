package db

import (
	"database/sql"
	"go-network-scanner/api"
)

type ConfigRepo struct {
	db      *sql.DB
	Current *api.Config
}

func (c *ConfigRepo) Get(name string) (api.Config, error) {
	return get(c.db, CONFIG_SELECT_BY_NAME, mapConfig, name)
}

func (c *ConfigRepo) All() ([]api.Config, error) {
	return all(
		c.db,
		CONFIG_SELECT_ALL,
		mapConfig,
	)
}

func (c *ConfigRepo) Insert(config api.Config) (int, error) {
	return upsert(
		c.db,
		CONFIG_INSERT,
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
