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
	return api.EventNmapScan
}

func (l *UiListener) Listen(i ...interface{}) {
	go l.process(i[0].(api.Scan))
}

func (l *UiListener) process(s api.Scan) {
	log.Debug("Scan received in UI listener", log.Any("scan", s))

	l.ui.AddScan(s)

	log.Debug("Scan displayed in UI listener", log.Any("scan", s))
}
