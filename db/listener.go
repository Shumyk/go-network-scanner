package db

import (
	"go-network-scanner/api"
	log "log/slog"
	"time"

	"github.com/kataras/go-events"
)

type DbListener struct {
	hosts *HostRepo
	scans *ScanRepo
}

func (l *DbListener) Topic() events.EventName {
	return "scan"
}

func (l *DbListener) Listen(i ...interface{}) {
	go process(i[0].(api.ScanFound))
}

func process(scan api.ScanFound) {
	log.Info("Scan received in DB listener", log.Any("scan", scan))
	time.Sleep(2 * time.Second)
	log.Info("Scan persisted in DB listener", log.Any("scan", scan))
	// TODO: implement persisting scan
}
