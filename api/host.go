package api

import "time"

type Host struct {
	Ip              string
	FirstDiscovered time.Time
	LastSeen        time.Time
	OsType          string
	HostnameLabel   string
	CustomLabel     string
}
