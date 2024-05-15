package pool

import (
	"go-network-scanner/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorkers(t *testing.T) {
	tests := []struct {
		name       string
		conf       *api.Config
		wantsCount int
	}{
		{
			name: "Test with 1 IP range",
			conf: &api.Config{
				IPRange1: "ip-range1",
			},
			wantsCount: 1,
		},
		{
			name: "Test with 3 IP ranges",
			conf: &api.Config{
				IPRange1: "ip-range1",
				IPRange2: "ip-range2",
				IPRange3: "ip-range3",
			},
			wantsCount: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workers := NewWorkers(tt.conf)
			assert.Equal(t, tt.wantsCount, len(workers))
		})
	}
}
