package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateTimeFrom(t *testing.T) {
	t.Run("should return a DateTime from a time.Time", func(t *testing.T) {
		// Given
		tt := time.Now()
		// When
		dt := DateTimeFrom(tt)
		// Then
		assert.Equal(t, tt, time.Time(dt))

		bytes, err := dt.MarshalText()
		assert.Nil(t, err)
		assert.Equal(t, tt.Format(time.DateTime), string(bytes))
	})
}

func TestMergeResults(t *testing.T) {
	t.Run("should merge multiple scanned results into one", func(t *testing.T) {
		// Given
		r1 := &ScannedResultsDTO{
			Data: []HostDTO{
				{IP: "ip1"},
				{IP: "ip2"},
			},
		}
		r2 := &ScannedResultsDTO{
			Data: []HostDTO{
				{IP: "ip2"},
				{},
			},
		}
		// When
		m := MergeResults([]*ScannedResultsDTO{r1, r2})
		// Then
		assert.Len(t, m.Data, 4)
	})
}

func TestHostDTOFrom(t *testing.T) {
	t.Run("should convert a Host to a HostDTO", func(t *testing.T) {
		// Given
		h := &Host{
			FirstDiscovered: time.Now(),
			LastSeen:        time.Now(),
			IP:              "ip",
			OsType:          "os",
			HostnameLabel:   "hostname",
			CustomLabel:     "custom",
		}
		// When
		dto := HostDTOFrom(h)
		// Then
		assert.Equal(t, h.IP, dto.IP)
		assert.Equal(t, h.OsType, dto.OS)
		assert.Equal(t, h.HostnameLabel, dto.HostnameLabel)
		assert.Equal(t, h.CustomLabel, dto.CustomLabel)
		assert.Len(t, dto.Services, 0)
		assert.Equal(t, h.FirstDiscovered, time.Time(dto.FirstDiscovered))
		assert.Equal(t, h.LastSeen, time.Time(dto.LastDiscovered))
	})
}

func TestScanDTOFrom(t *testing.T) {
	t.Run("should convert a Scan to a ScanDTO", func(t *testing.T) {
		// Given
		s := &Scan{
			FirstDiscovered: time.Now(),
			LastSeen:        time.Now(),
			IP:              "ip",
			Port:            80,
			ServiceBanner:   "http - nginx",
		}
		// When
		dto := ScanDTOFrom(s)
		// Then
		assert.Equal(t, s.IP, dto.IP)
		assert.Equal(t, s.Port, dto.Port)
		assert.Equal(t, s.ServiceBanner, dto.Banner)
		assert.Equal(t, s.FirstDiscovered, time.Time(dto.FirstDiscovered))
		assert.Equal(t, s.LastSeen, time.Time(dto.LastDiscovered))
	})
}
