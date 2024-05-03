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
	log.Info("Scan received in DB listener", log.Any("scan", found))

	// Persist the scan in the database
	scan := api.Scan{
		Ip:            found.Ip,
		Port:          found.Port,
		ServiceBanner: found.ServiceBanner,
	}
	id, err := l.scans.Insert(scan)
	if err != nil {
		log.Error("Failed to persist scan in DB", log.Any("error", err))
		return
	}
	log.Info(
		"Scan persisted in DB listener",
		log.Any("scan.ip", scan.Id),
		log.Int("id", id),
	)
}
