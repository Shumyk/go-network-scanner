package api

import "time"

// DateTime is a custom type for time.Time
// that serializes to a string in the format "2006-01-02 15:04:05".
type DateTime time.Time

// MarshalText serializes a DateTime to a string.
func (dt DateTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(dt).Format(time.DateTime)), nil
}

// DateTimeFrom creates a DateTime from a time.Time.
func DateTimeFrom(t time.Time) DateTime {
	return DateTime(t)
}

// ScannedResultsDTO is the data transfer object for scanned results.
type ScannedResultsDTO struct {
	Data []HostDTO `json:"data"`
}

// MergeResults merges multiple scanned results into one.
func MergeResults(results []*ScannedResultsDTO) *ScannedResultsDTO {
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

// HostDTOFrom converts a Host to a HostDTO.
func HostDTOFrom(h *Host) *HostDTO {
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
func ScanDTOFrom(s *Scan) *ScanDTO {
	return &ScanDTO{
		FirstDiscovered: DateTimeFrom(s.FirstDiscovered),
		LastDiscovered:  DateTimeFrom(s.LastSeen),

		IP:     s.IP,
		Banner: s.ServiceBanner,
		Port:   s.Port,
	}
}
