package database

import (
	"database/sql"
	"time"
)

type Configs struct {
	db *sql.DB
}

type Config struct {
	name        string
	enableProxy bool
	ipRange1    string
	ipRange2    string
	ipRange3    string
	createdAt   time.Time
}

func (c *Configs) Get(name string) (Config, error) {
	var config Config
	err := db.
		QueryRow(CONFIG_SELECT_BY_NAME, name).
		Scan(&config.name, &config.enableProxy, &config.ipRange1, &config.ipRange2, &config.ipRange3)
	if err != nil {
		return Config{}, err
	}
	return config, nil

}

func (c *Configs) All() ([]Config, error) {
	rows, err := db.Query(CONFIG_SELECT_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []Config
	for rows.Next() {
		var config Config
		err := rows.Scan(&config.name, &config.enableProxy, &config.ipRange1, &config.ipRange2, &config.ipRange3)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func (c *Configs) Insert(config Config) (int, error) {
	result, err := db.Exec(CONFIG_INSERT, config.name, config.enableProxy, config.ipRange1, config.ipRange2, config.ipRange3)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
