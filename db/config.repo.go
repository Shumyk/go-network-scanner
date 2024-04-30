package db

import (
	"go-network-scanner/api"
)

type ConfigRepo struct {
	// DB *sql.DB
	Current api.Config
}

func NewConfigRepo(current api.Config) *ConfigRepo {
	repo := &ConfigRepo{Current: current}
	repo.Insert(current)
	return repo
}

func (c *ConfigRepo) Get(name string) (api.Config, error) {
	var config api.Config
	err := db.
		QueryRow(CONFIG_SELECT_BY_NAME, name).
		Scan(&config.Name, &config.EnableProxy, &config.IpRange1, &config.IpRange2, &config.IpRange3)
	if err != nil {
		return api.Config{}, err
	}
	return config, nil

}

func (c *ConfigRepo) All() ([]api.Config, error) {
	rows, err := db.Query(CONFIG_SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []api.Config
	for rows.Next() {
		var config api.Config
		err := rows.Scan(&config.Name, &config.EnableProxy, &config.IpRange1, &config.IpRange2, &config.IpRange3)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func (c *ConfigRepo) Insert(config api.Config) (int, error) {
	result, err := db.Exec(CONFIG_INSERT, config.Name, config.EnableProxy, config.IpRange1, config.IpRange2, config.IpRange3)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
