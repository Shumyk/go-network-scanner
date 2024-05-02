package api

import "time"

type Scan struct {
	Id              int
	Ip              string
	Port            uint16
	FirstDiscovered time.Time
	LastSeen        time.Time
	ServiceBanner   string
}

type ScanFound struct {
	Ip            string
	Port          uint16
	ServiceBanner string
	OsType        string
	HostnameLabel string
	CustomLabel   string
}
