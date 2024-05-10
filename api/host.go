package api

import "time"

// Host represents a host entry.
type Host struct {
	Session         string
	IP              string
	FirstDiscovered time.Time
	LastSeen        time.Time
	OsType          string
	HostnameLabel   string
	CustomLabel     string
}
