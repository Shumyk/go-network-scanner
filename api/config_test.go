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
	assert.Empty(t, c.CreatedAt)
}

func TestIPRanges(t *testing.T) {
	tests := []struct {
		name      string
		ip1       string
		ip2       string
		ip3       string
		want      []string
		wantCount int
	}{
		{name: "all empty", ip1: "", ip2: "", ip3: "", want: []string{}},
		{
			name: "all filled", ip1: "192.43.56.32", ip2: "192.43.56.35", ip3: "192.43.56.38",
			want: []string{"192.43.56.32", "192.43.56.35", "192.43.56.38"},
		},
		{
			name: "some empty",
			ip1:  "192.43.56.38", ip3: "192.43.56.33",
			want: []string{"192.43.56.38", "192.43.56.33"},
		},
		{name: "only one", ip1: "192.43.56.38", want: []string{"192.43.56.38"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{IPRange1: tt.ip1, IPRange2: tt.ip2, IPRange3: tt.ip3}
			assert.Equal(t, tt.want, c.IPRanges())
			assert.Equal(t, len(tt.want), c.IPRangesCount())
		})
	}
}
