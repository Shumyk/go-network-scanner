package ui

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

type UiListener struct{}

func (l *UiListener) Topic() events.EventName {
	return "scan"
}

func (l *UiListener) Listen(i ...interface{}) {
	scan := i[0].(api.ScanFound)
	log.Info("Scan received in UI listener", log.Any("scan", scan))
}
