package ui

import (
	"go-network-scanner/api"
	log "log/slog"
	"time"

	"github.com/kataras/go-events"
)

type UiListener struct{}

func (l *UiListener) Topic() events.EventName {
	return api.EVENT_SCAN
}

func (l *UiListener) Listen(i ...interface{}) {
	go process(i[0].(api.Scan))
}

func process(s api.Scan) {
	log.Debug("Scan received in UI listener", log.Any("scan", s))
	time.Sleep(2 * time.Second)
	log.Debug("Scan displayed in UI listener", log.Any("scan", s))
}
