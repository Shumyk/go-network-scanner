package db

const (
	// The configuration entered
	CREATE_TABLE_CONFIG = `
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
	CREATE_TABLE_SCAN = `
		CREATE TABLE IF NOT EXISTS scan
		(
			session          TEXT     NOT NULL,
			ip               TEXT     NOT NULL,
			port             INTEGER  NOT NULL,
			first_discovered DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			last_seen        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			service_banner   TEXT     NOT NULL,
			PRIMARY KEY (session, ip, port)
		);
	`
	// The unique ips found (ip, time of first discovery, time when it stopped being found, operating system type, hostname label and custom label)
	CREATE_TABLE_HOST = `
		CREATE TABLE IF NOT EXISTS host
		(
			session          TEXT     NOT NULL,
			ip               TEXT     NOT NULL,
			first_discovered DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			last_seen        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			os_type          TEXT     NOT NULL,
			hostname_label   TEXT     NOT NULL,
			custom_label     TEXT     NOT NULL,

			PRIMARY KEY (session, ip)
		);
	`

	CONFIG_SELECT_ALL     = "SELECT name, enable_proxy, ip_range1, ip_range2, ip_range3 FROM config;"
	CONFIG_SELECT_BY_NAME = "SELECT name, enable_proxy, ip_range1, ip_range2, ip_range3 FROM config WHERE name = ?;"
	CONFIG_INSERT         = "INSERT INTO config (name, enable_proxy, ip_range1, ip_range2, ip_range3) VALUES (?, ?, ?, ?, ?);"

	SCAN_SELECT_ALL = "SELECT id, ip, port, first_discovered, last_seen, service_banner FROM scan;"
	SCAN_INSERT     = "INSERT INTO scan (ip, port, last_seen, service_banner) VALUES (?, ?, datetime(), ?);"
	SCAN_UPSERT     = `
		INSERT INTO scan (session, ip, port, last_seen, service_banner) VALUES (:session, :ip, :port, datetime(), :serviceBanner)
			ON CONFLICT(session, ip, port) DO UPDATE SET last_seen = datetime(), service_banner = :serviceBanner;
	`

	HOST_SELECT_ALL = "SELECT ip, first_discovered, last_seen, os_type, hostname_label, custom_label FROM host;"
	HOST_INSERT     = "INSERT INTO host (ip, last_seen, os_type, hostname_label, custom_label) VALUES (?, datetime(), ?, ?, ?);"
	HOST_UPSERT     = `
	INSERT INTO host (session, ip, last_seen, os_type, hostname_label, custom_label) VALUES (:session, :ip, datetime(), :osType, :hostnameLabel, :customLabel)
 	   ON CONFLICT(session, ip) DO UPDATE SET last_seen = datetime(), os_type = :osType, hostname_label = :hostnameLabel, custom_label = :customLabel;
	`
)
