package api

import "time"

type Config struct {
	Name        string
	EnableProxy bool
	IpRange1    string
	IpRange2    string
	IpRange3    string
	CreatedAt   time.Time
}
