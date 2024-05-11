package db

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

// ScanListener listens for scan events and persists them.
type ScanListener struct {
	scans *ScanRepo
}

// Topic returns the event topic.
func (*ScanListener) Topic() events.EventName {
	return api.EventNmapScan
}

// Listen listens for scan events.
func (l *ScanListener) Listen(i ...any) {
	s, ok := i[zero].(api.Scan)
	if !ok {
		log.Error("Failed to cast to scan")
		return
	}
	go l.process(&s)
}

func (l *ScanListener) process(scan *api.Scan) {
	log.Debug("Scan received", log.Any("scan", scan))

	id, err := l.scans.Upsert(scan)
	if err != nil {
		log.Error("Failed to persist scan", log.Any("error", err))
		return
	}

	log.Debug(
		"Scan persisted",
		log.Any("ip", scan.IP),
		log.Any("port", scan.Port),
		log.Int("id", id),
	)
}
