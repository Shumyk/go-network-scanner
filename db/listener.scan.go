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
	return "scan"
}

func (l *ScanListener) Listen(i ...interface{}) {
	go l.process(i[0].(api.ScanFound))
}

func (l *ScanListener) process(found api.ScanFound) {
	log.Info("Scan received", log.Any("scan", found))

	scan := api.ScanOf(found)
	id, err := l.scans.Upsert(scan)
	if err != nil {
		log.Error("Failed to persist scan", log.Any("error", err))
		return
	}
	log.Info(
		"Scan persisted",
		log.Any("ip", scan.Ip),
		log.Any("port", scan.Port),
		log.Int("id", id),
	)
}
