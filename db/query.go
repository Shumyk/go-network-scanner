package db

const (
	// The configuration entered
	create_table_config = `
		CREATE TABLE IF NOT EXISTS config
		(
			name         TEXT    PRIMARY KEY,
			enable_proxy BOOLEAN NOT NULL,
			ip_range1    TEXT    NOT NULL,
			ip_range2    TEXT,
			ip_range3    TEXT,
			created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`
	// scan (ip, port, time if first discovery, time when it stopped being found, service banner)
	create_table_scan = `
		CREATE TABLE IF NOT EXISTS scan
		(
			id               INTEGER PRIMARY KEY,
			ip               TEXT     NOT NULL,
			port             INTEGER  NOT NULL,
			first_discovered DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			last_seen        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			service_banner   TEXT     NOT NULL,
			created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`
	// The unique ips found (ip, time of first discovery, time when it stopped being found, operating system type, hostname label and custom label)
	create_table_host = `
		CREATE TABLE IF NOT EXISTS host
		(
			id               INTEGER PRIMARY KEY,
			ip               TEXT     NOT NULL,
			first_discovered DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			last_seen        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			os_type          TEXT     NOT NULL,
			hostname_label   TEXT     NOT NULL,
			custom_label     TEXT     NOT NULL,
			created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	CONFIG_SELECT_ALL     = "SELECT name, enable_proxy, ip_range1, ip_range2, ip_range3 FROM config;"
	CONFIG_SELECT_BY_NAME = "SELECT name, enable_proxy, ip_range1, ip_range2, ip_range3 FROM config WHERE name = ?;"
	CONFIG_INSERT         = "INSERT INTO config (name, enable_proxy, ip_range1, ip_range2, ip_range3) VALUES (?, ?, ?, ?, ?);"
)
