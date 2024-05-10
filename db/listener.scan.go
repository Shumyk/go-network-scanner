package db

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

type ScanListener struct {
	scans *ScanRepo
}

func (l *ScanListener) Topic() events.EventName {
	return api.EventNmapScan
}

func (l *ScanListener) Listen(i ...interface{}) {
	go l.process(i[0].(api.Scan))
}

func (l *ScanListener) process(scan api.Scan) {
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
