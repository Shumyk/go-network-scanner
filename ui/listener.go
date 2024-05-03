package ui

import (
	"go-network-scanner/api"
	log "log/slog"
	"time"

	"github.com/kataras/go-events"
)

type UiListener struct{}

func (l *UiListener) Topic() events.EventName {
	return "scan"
}

func (l *UiListener) Listen(i ...interface{}) {
	go process(i[0].(api.ScanFound))
}

func process(s api.ScanFound) {
	log.Info("Scan received in UI listener", log.Any("scan", s))
	time.Sleep(2 * time.Second)
	log.Info("Scan displayed in UI listener", log.Any("scan", s))
}
