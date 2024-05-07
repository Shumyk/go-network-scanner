package ui

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

type UiListener struct {
	ui *Ui
}

func (l *UiListener) Topic() events.EventName {
	return api.EVENT_SCAN
}

func (l *UiListener) Listen(i ...interface{}) {
	go l.process(i[0].(api.Scan))
}

func (l *UiListener) process(s api.Scan) {
	log.Info("Scan received in UI listener", log.Any("scan", s))

	l.ui.AddScan(s)

	log.Info("Scan displayed in UI listener", log.Any("scan", s))
}
