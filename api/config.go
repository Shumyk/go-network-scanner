// Package api provides common models and functions.
package api

import (
	"slices"
	"time"
)

const empStr = ""

// ConfDTO represents a configuration data transfer object sent from UI.
type ConfDTO struct {
	DatabaseName string
	SessionName  string
	IPRange1     string
	IPRange2     string
	IPRange3     string
	Proxy1       string
	Proxy2       string
	Proxy3       string
	EnableProxy  bool
}

// Config represents a configuration entry.
type Config struct {
	CreatedAt   time.Time
	Name        string
	IPRange1    string
	Proxy1      string
	IPRange2    string
	Proxy2      string
	IPRange3    string
	Proxy3      string
	EnableProxy bool
}

// EmptyConfig returns a Config with empty fields.
func EmptyConfig() Config {
	return Config{
		Name:        empStr,
		EnableProxy: false,
		IPRange1:    empStr,
		IPRange2:    empStr,
		IPRange3:    empStr,
		CreatedAt:   time.Time{},
	}
}

// IPRanges returns the non-empty IP ranges.
func (c *Config) IPRanges() []string {
	return slices.DeleteFunc(
		[]string{c.IPRange1, c.IPRange2, c.IPRange3},
		func(ip string) bool {
			return ip == empStr
		},
	)
}

// IPRangesCount returns the number of non-empty IP ranges.
func (c *Config) IPRangesCount() int {
	return len(c.IPRanges())
}
