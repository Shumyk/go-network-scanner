package api

import "time"

type Host struct {
	Session         string
	Ip              string
	FirstDiscovered time.Time
	LastSeen        time.Time
	OsType          string
	HostnameLabel   string
	CustomLabel     string
}
