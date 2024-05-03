INSERT INTO scan (ip, port, last_seen, service_banner) VALUES (?, ?, datetime(), ?);

INSERT INTO scan (ip, port, last_seen, service_banner) VALUES (?, ?, datetime(), ?)
    ON CONFLICT(ip, port) DO UPDATE SET last_seen = datetime(), service_banner = ?;



SELECT ip, first_discovered, last_seen, os_type, hostname_label, custom_label FROM host;

INSERT INTO host (ip, last_seen, os_type, hostname_label, custom_label) VALUES (?, datetime(), ?, ?, ?);

INSERT INTO host (ip, last_seen, os_type, hostname_label, custom_label) VALUES (?, datetime(), ?, ?, ?)
    ON CONFLICT(ip) DO UPDATE SET last_seen = datetime(), os_type = ?, hostname_label = ?, custom_label = ?;