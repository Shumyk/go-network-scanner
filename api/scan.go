package api

import "time"

type Scan struct {
	Id              int
	Ip              string
	Port            int
	FirstDiscovered time.Time
	LastSeen        time.Time
	ServiceBanner   string
}
