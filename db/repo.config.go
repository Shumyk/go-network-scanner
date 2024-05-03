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
	return Get(c.db, CONFIG_SELECT_BY_NAME, mapConfig, name)
}

func (c *ConfigRepo) All() ([]api.Config, error) {
	return All(
		c.db,
		CONFIG_SELECT_ALL,
		mapConfig,
	)
}

func (c *ConfigRepo) Insert(config api.Config) (int, error) {
	return Upsert(
		c.db,
		CONFIG_INSERT,
		config.Name, config.EnableProxy,
		config.IpRange1, config.IpRange2, config.IpRange3,
	)
}

func mapConfig(f func(...any) error) (api.Config, error) {
	var c api.Config
	err := f(
		&c.Name,
		&c.EnableProxy,
		&c.IpRange1,
		&c.IpRange2,
		&c.IpRange3,
	)
	return c, err
}
