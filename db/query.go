//nolint:revive // Ignore lengthy queries
package db

const (
	// The configuration entered.
	queryCreateTableConfig = `
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
	queryCreateTableScan = `
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
	queryCreateTableHost = `
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

	queryConfigSelectAll    = "SELECT name, enable_proxy, ip_range1, ip_range2, ip_range3 FROM config;"
	queryConfigSelectByName = "SELECT name, enable_proxy, ip_range1, ip_range2, ip_range3 FROM config WHERE name = ?;"
	queryConfigInsert       = "INSERT INTO config (name, enable_proxy, ip_range1, ip_range2, ip_range3) VALUES (?, ?, ?, ?, ?);"
	queryConfigUpsert       = `
		INSERT INTO config (name, enable_proxy, ip_range1, ip_range2, ip_range3) VALUES (:name, :enableProxy, :ipRange1, :ipRange2, :ipRange3)
			ON CONFLICT(name) DO UPDATE SET enable_proxy = :enableProxy, ip_range1 = :ipRange1, ip_range2 = :ipRange2, ip_range3 = :ipRange3;
	`

	queryScanSelectAll = "SELECT id, ip, port, first_discovered, last_seen, service_banner FROM scan;"
	queryScanInsert    = "INSERT INTO scan (ip, port, last_seen, service_banner) VALUES (?, ?, datetime(), ?);"
	queryScanUpsert    = `
		INSERT INTO scan (session, ip, port, last_seen, service_banner) VALUES (:session, :ip, :port, datetime(), :serviceBanner)
			ON CONFLICT(session, ip, port) DO UPDATE SET last_seen = datetime(), service_banner = :serviceBanner;
	`

	queryHostSelectAll = "SELECT ip, first_discovered, last_seen, os_type, hostname_label, custom_label FROM host;"
	queryHostInsert    = "INSERT INTO host (ip, last_seen, os_type, hostname_label, custom_label) VALUES (?, datetime(), ?, ?, ?);"
	queryHostUpsert    = `
	INSERT INTO host (session, ip, last_seen, os_type, hostname_label, custom_label) VALUES (:session, :ip, datetime(), :osType, :hostnameLabel, :customLabel)
 	   ON CONFLICT(session, ip) DO UPDATE SET last_seen = datetime(), os_type = :osType, hostname_label = :hostnameLabel, custom_label = :customLabel;
	`
)
