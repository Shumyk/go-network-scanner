package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyConfig(t *testing.T) {
	c := EmptyConfig()
	assert.Empty(t, c.Name)
	assert.False(t, c.EnableProxy)
	assert.Empty(t, c.IPRange1)
	assert.Empty(t, c.IPRange2)
	assert.Empty(t, c.IPRange3)
	assert.Empty(t, c.Proxy1)
	assert.Empty(t, c.Proxy2)
	assert.Empty(t, c.Proxy3)
	assert.Empty(t, c.CreatedAt)
}

func TestConfigOf(t *testing.T) {
	dto := &ConfDTO{
		DatabaseName: "database name",
		SessionName:  "session ",
		IPRange1:     " hostname.com ",
		IPRange2:     "192.24.43.23 ",
		IPRange3:     "  ",
		Proxy1:       "pr oxy1",
		Proxy2:       "proxy 2",
		Proxy3:       "Proxy3",
		EnableProxy:  true,
	}
	c := ConfigOf(dto)
	assert.Equal(t, "session", c.Name)
	assert.True(t, c.EnableProxy)
	assert.Equal(t, "hostname.com", c.IPRange1)
	assert.Equal(t, "192.24.43.23", c.IPRange2)
	assert.Empty(t, c.IPRange3)
	assert.Equal(t, "pr oxy1", c.Proxy1)
	assert.Equal(t, "proxy 2", c.Proxy2)
	assert.Equal(t, "Proxy3", c.Proxy3)
	assert.NotNil(t, c.CreatedAt)
}

func TestIPRanges(t *testing.T) {
	tests := []struct {
		name        string
		ip1         string
		ip2         string
		ip3         string
		enableProxy bool
		proxy1      string
		proxy2      string
		proxy3      string
		want        []*IPRange
	}{
		{name: "all empty", ip1: "", ip2: "", ip3: "", want: []*IPRange{}},
		{
			name:        "all filled",
			enableProxy: true,
			ip1:         "192.43.56.32", proxy1: "proxy1",
			ip2: "192.43.56.35",
			ip3: "192.43.56.39", proxy3: "proxy3",
			want: []*IPRange{
				{IPRange: "192.43.56.32", Proxy: "proxy1"},
				{IPRange: "192.43.56.35"},
				{IPRange: "192.43.56.39", Proxy: "proxy3"},
			},
		},
		{
			name:        "some empty",
			enableProxy: false,
			ip1:         "192.43.56.45", proxy1: "proxy1",
			ip3: "192.43.56.33", proxy3: "proxy3",
			want: []*IPRange{
				{IPRange: "192.43.56.45"},
				{IPRange: "192.43.56.33"},
			},
		},
		{
			name:        "only one with proxy",
			enableProxy: true,
			ip1:         "192.43.56.89", proxy1: "proxy1",
			proxy2: "proxy2",
			proxy3: "proxy3",
			want:   []*IPRange{{IPRange: "192.43.56.89", Proxy: "proxy1"}},
		},
		{
			name:        "only one with disabled proxy",
			enableProxy: false,
			ip1:         "192.43.56.38", proxy1: "proxy1",
			proxy2: "proxy2",
			proxy3: "proxy3",
			want:   []*IPRange{{IPRange: "192.43.56.38"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				EnableProxy: tt.enableProxy,
				IPRange1:    tt.ip1, Proxy1: tt.proxy1,
				IPRange2: tt.ip2, Proxy2: tt.proxy2,
				IPRange3: tt.ip3, Proxy3: tt.proxy3,
			}
			assert.Equal(t, tt.want, c.IPRanges())
			assert.Equal(t, len(tt.want), c.IPRangesCount())
		})
	}
}
