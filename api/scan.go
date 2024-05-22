package api

import "time"

// Scan represents a scan entry.
type Scan struct {
	FirstDiscovered time.Time
	LastSeen        time.Time

	Session       string
	IP            string
	ServiceBanner string
	Port          uint16
}
