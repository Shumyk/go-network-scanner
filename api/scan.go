package api

import "time"

type Scan struct {
	Session         string
	Ip              string
	Port            uint16
	FirstDiscovered time.Time
	LastSeen        time.Time
	ServiceBanner   string
}
