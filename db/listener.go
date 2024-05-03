package db

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

type DbListener struct {
	hosts HostRepo
	scans ScanRepo
}

func (l *DbListener) Topic() events.EventName {
	return "scan"
}

func (l *DbListener) Listen(i ...interface{}) {
	scan := i[0].(api.ScanFound)
	log.Info("Scan received", log.Any("scan", scan))
}
