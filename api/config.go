package api

import (
	"slices"
	"time"
)

type Config struct {
	Name        string
	EnableProxy bool
	IpRange1    string
	IpRange2    string
	IpRange3    string
	CreatedAt   time.Time
}

func (c *Config) IpRanges() []string {
	return slices.DeleteFunc(
		[]string{c.IpRange1, c.IpRange2, c.IpRange3},
		func(ip string) bool {
			return len(ip) == 0
		},
	)
}

func (c *Config) IpRangesCount() int {
	return len(c.IpRanges())
}
