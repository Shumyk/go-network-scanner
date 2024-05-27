package api

import "time"

type DateTime time.Time

func (dt DateTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(dt).Format(time.DateTime)), nil
}

// ScannedResultsDTO is the data transfer object for scanned results.
type ScannedResultsDTO struct {
	Data []HostDTO `json:"data"`
}

// MergeResults merges multiple scanned results into one.
func MergeResults(results []*ScannedResultsDTO) *ScannedResultsDTO {
	// Merge results
	m := &ScannedResultsDTO{
		Data: []HostDTO{},
	}
	for _, r := range results {
		m.Data = append(m.Data, r.Data...)
	}
	return m
}

// HostDTO is the data transfer object for scanned results.
type HostDTO struct {
	FirstDiscovered DateTime `json:"firstDiscovered"`
	LastDiscovered  DateTime `json:"lastDiscovered"`

	IP            string `json:"ip"`
	OS            string `json:"os"`
	HostnameLabel string `json:"hostnameLabel"`
	CustomLabel   string `json:"customLabel"`

	Services []ScanDTO `json:"services"`
}

func DateTimeFrom(t time.Time) DateTime {
	return DateTime(t)
}

// HostDTOFrom converts a Host to a HostDTO.
func HostDTOFrom(h *Host, scannedAt time.Time) *HostDTO {
	return &HostDTO{
		FirstDiscovered: DateTimeFrom(h.FirstDiscovered),
		LastDiscovered:  DateTimeFrom(h.LastSeen),

		IP:            h.IP,
		OS:            h.OsType,
		HostnameLabel: h.HostnameLabel,
		CustomLabel:   h.CustomLabel,

		Services: []ScanDTO{},
	}
}

// ScanDTO is the data transfer object for scans.
type ScanDTO struct {
	FirstDiscovered DateTime `json:"firstDiscovered"`
	LastDiscovered  DateTime `json:"lastDiscovered"`

	IP     string `json:"ip"`
	Banner string `json:"banner"`
	Port   uint16 `json:"port"`
}

// ScanDTOFrom creates a ScanDTO from a Scan.
func ScanDTOFrom(s *Scan, scannedAt time.Time) *ScanDTO {
	return &ScanDTO{
		FirstDiscovered: DateTimeFrom(s.FirstDiscovered),
		LastDiscovered:  DateTimeFrom(s.LastSeen),

		IP:     s.IP,
		Banner: s.ServiceBanner,
		Port:   s.Port,
	}
}
