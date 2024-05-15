package nmap

import (
	"testing"

	"github.com/Ullaakut/nmap/v3"
	"github.com/stretchr/testify/assert"
)

const ip = "127.0.0.1"

func TestIP(t *testing.T) {
	tests := []struct {
		name string
		h    *nmap.Host
		want string
	}{
		{
			name: "empty",
			h:    &nmap.Host{},
			want: "",
		},
		{
			name: "one address",
			h: &nmap.Host{
				Addresses: []nmap.Address{
					{Addr: ip},
				},
			},
			want: ip,
		},
		{
			name: "two addresses",
			h: &nmap.Host{
				Addresses: []nmap.Address{
					{Addr: ip},
					{Addr: ip + "1"},
				},
			},
			want: ip,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IP(tt.h))
		})
	}
}

func TestHostname(t *testing.T) {
	tests := []struct {
		name string
		h    *nmap.Host
		want string
	}{
		{
			name: "empty",
			h:    &nmap.Host{},
			want: "",
		},
		{
			name: "one hostname",
			h: &nmap.Host{
				Hostnames: []nmap.Hostname{
					{Name: "localhost"},
				},
			},
			want: "localhost",
		},
		{
			name: "two hostnames",
			h: &nmap.Host{
				Hostnames: []nmap.Hostname{
					{Name: "localhost"},
					{Name: "localhost.local"},
				},
			},
			want: "localhost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Hostname(tt.h))
		})
	}
}

func TestOS(t *testing.T) {
	tests := []struct {
		name string
		h    *nmap.Host
		want string
	}{
		{
			name: "empty",
			h:    &nmap.Host{},
			want: "",
		},
		{
			name: "one os",
			h: &nmap.Host{
				OS: nmap.OS{
					Matches: []nmap.OSMatch{
						{Name: "Linux"},
					},
				},
			},
			want: "Linux",
		},
		{
			name: "two os",
			h: &nmap.Host{
				OS: nmap.OS{
					Matches: []nmap.OSMatch{
						{Name: "Linux"},
						{Name: "Windows"},
					},
				},
			},
			want: "Linux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OS(tt.h))
		})
	}
}

func TestServiceBanner(t *testing.T) {
	tests := []struct {
		name string
		p    *nmap.Port
		want string
	}{
		{
			name: "empty",
			p:    &nmap.Port{},
			want: "",
		},
		{
			name: "all fields",
			p: &nmap.Port{
				Service: nmap.Service{
					Name:    "http",
					Product: "nginx",
					Version: "1.0",
				},
			},
			want: "http - nginx - 1.0",
		},
		{
			name: "name only",
			p: &nmap.Port{
				Service: nmap.Service{
					Name: "http",
				},
			},
			want: "http",
		},
		{
			name: "name and product",
			p: &nmap.Port{
				Service: nmap.Service{
					Name:    "http",
					Product: "nginx",
				},
			},
			want: "http - nginx",
		},
		{
			name: "name and version",
			p: &nmap.Port{
				Service: nmap.Service{
					Name:    "http",
					Version: "1.0",
				},
			},
			want: "http - 1.0",
		},
		{
			name: "product only",
			p: &nmap.Port{
				Service: nmap.Service{
					Product: "nginx",
				},
			},
			want: "nginx",
		},
		{
			name: "version only",
			p: &nmap.Port{
				Service: nmap.Service{
					Version: "1.0",
				},
			},
			want: "1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ServiceBanner(tt.p))
		})
	}
}
