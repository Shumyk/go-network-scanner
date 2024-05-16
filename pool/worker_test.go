package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/util"
	"testing"

	"github.com/Ullaakut/nmap/v3"
	"github.com/kataras/go-events"
	"github.com/stretchr/testify/assert"
)

type scannerMock struct {
	returns *nmap.Run
}

func (s *scannerMock) Scan() *nmap.Run {
	return s.returns
}

func TestDo(t *testing.T) {
	tests := []struct {
		name            string
		returns         *nmap.Run
		wantsHostEvents int
		wantsScanEvents int
	}{
		{
			name: "Test with 1 host and 1 scan",
			returns: &nmap.Run{
				Hosts: []nmap.Host{
					{
						Addresses: []nmap.Address{
							{
								Addr: "localhost",
							},
						},
						Ports: []nmap.Port{
							{
								ID: 80,
							},
						},
					},
				},
			},
			wantsHostEvents: 1,
			wantsScanEvents: 1,
		},
		{
			name: "Test with 2 hosts and 2 scans",
			returns: &nmap.Run{
				Hosts: []nmap.Host{
					{
						Addresses: []nmap.Address{
							{
								Addr: "localhost",
							},
						},
						Ports: []nmap.Port{
							{
								ID: 80,
							},
						},
					},
					{
						Addresses: []nmap.Address{
							{
								Addr: "localhost",
							},
						},
						Ports: []nmap.Port{
							{
								ID: 80,
							},
						},
					},
				},
			},
			wantsHostEvents: 2,
			wantsScanEvents: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events.On(api.EventNmapHost, func(e ...any) {
				tt.wantsHostEvents--
			})
			events.On(api.EventNmapScan, func(e ...any) {
				tt.wantsScanEvents--
			})

			w := &worker{
				nmap: &scannerMock{tt.returns},
				log:  util.Logger(),
			}
			w.Do()

			assert.Zero(t, tt.wantsHostEvents)
			assert.Zero(t, tt.wantsScanEvents)
		})
	}
}
