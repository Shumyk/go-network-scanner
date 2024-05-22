package api

import "time"

// ScannedResultsDTO is the data transfer object for scanned results.
type ScannedResultsDTO struct {
	Data []HostDTO `json:"data"`
}

// HostDTO is the data transfer object for scanned results.
type HostDTO struct {
	FirstDiscovered time.Time `json:"firstDiscovered"`
	LastDiscovered  time.Time `json:"lastDiscovered"`

	IP            string `json:"ip"`
	OS            string `json:"os"`
	HostnameLabel string `json:"hostnameLabel"`
	CustomLabel   string `json:"customLabel"`

	Services []ScanDTO `json:"services"`
}

// HostDTOFrom converts a Host to a HostDTO.
func HostDTOFrom(h *Host) *HostDTO {
	return &HostDTO{
		FirstDiscovered: h.FirstDiscovered,
		LastDiscovered:  h.LastSeen,

		IP:            h.IP,
		OS:            h.OsType,
		HostnameLabel: h.HostnameLabel,
		CustomLabel:   h.CustomLabel,
	}
}

// ScanDTO is the data transfer object for scans.
type ScanDTO struct {
	FirstDiscovered time.Time `json:"firstDiscovered"`
	LastDiscovered  time.Time `json:"lastDiscovered"`

	IP     string `json:"ip"`
	Banner string `json:"banner"`
	Port   uint16 `json:"port"`
}

// ScanDTOFrom creates a ScanDTO from a Scan.
func ScanDTOFrom(s *Scan) *ScanDTO {
	return &ScanDTO{
		FirstDiscovered: s.FirstDiscovered,
		LastDiscovered:  s.LastSeen,

		IP:     s.IP,
		Banner: s.ServiceBanner,
		Port:   s.Port,
	}
}
