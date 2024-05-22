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

// ScanDTO is the data transfer object for scans.
type ScanDTO struct {
	IP      string `json:"ip"`
	Service string `json:"service"`
	Port    uint16 `json:"port"`
}

// ScansDTO is the data transfer object for multiple scans by IP.
type ScansDTO struct {
	Data []*ScanDTO `json:"data"`
}

func ScanDTOFrom(s *Scan) *ScanDTO {
	return &ScanDTO{
		IP:      s.IP,
		Service: s.ServiceBanner,
		Port:    s.Port,
	}
}
