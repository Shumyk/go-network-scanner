// Package api provides common models and functions.
package api

import (
	"slices"
	"time"
)

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
	IPRange2    string
	IPRange3    string
	Proxy1      string
	Proxy2      string
	Proxy3      string
	EnableProxy bool
}

// EmptyConfig returns a Config with empty fields.
func EmptyConfig() Config {
	return Config{
		Name:        "",
		EnableProxy: false,
		IPRange1:    "",
		IPRange2:    "",
		IPRange3:    "",
		Proxy1:      "",
		Proxy2:      "",
		Proxy3:      "",
		CreatedAt:   time.Time{},
	}
}

// ConfigOf converts a ConfDTO to a Config.
func ConfigOf(dto *ConfDTO) *Config {
	return &Config{
		CreatedAt:   time.Now(),
		Name:        dto.SessionName,
		IPRange1:    dto.IPRange1,
		IPRange2:    dto.IPRange2,
		IPRange3:    dto.IPRange3,
		EnableProxy: dto.EnableProxy,
		Proxy1:      dto.Proxy1,
		Proxy2:      dto.Proxy2,
		Proxy3:      dto.Proxy3,
	}
}

// IPRange represents an IP range and its proxy.
type IPRange struct {
	IPRange string
	Proxy   string
}

// IPRanges returns the non-empty IP ranges.
func (c *Config) IPRanges() []*IPRange {
	return slices.DeleteFunc(
		[]*IPRange{
			{c.IPRange1, c.Proxy1},
			{c.IPRange2, c.Proxy2},
			{c.IPRange3, c.Proxy3},
		},
		func(ip *IPRange) bool {
			if !c.EnableProxy {
				//nolint:govet // It is used but not here.
				ip.Proxy = ""
			}
			return ip.IPRange == ""
		},
	)
}

// IPRangesCount returns the number of non-empty IP ranges.
func (c *Config) IPRangesCount() int {
	return len(c.IPRanges())
}
