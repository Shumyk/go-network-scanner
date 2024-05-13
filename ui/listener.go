package ui

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

const (
	zero  = 0
	topic = "topic"
	data  = "data"
)

type UiListener struct {
	ui *Ui
}

func (l *UiListener) Topic() events.EventName {
	return api.EventNmapScan
}

func (l *UiListener) Listen(i ...any) {
	if len(i) == zero || i[zero] == nil {
		log.Error(
			"no event data",
			log.Any(topic, l.Topic()),
		)
		return
	}
	t, ok := i[zero].(*api.Scan)
	if !ok {
		log.Error(
			"failed to cast",
			log.Any(topic, l.Topic()),
			log.Any(data, i[zero]),
		)
		return
	}
	//nolint:errcheck // Ignore error here as it's returned for tests.
	go l.process(t)
}

func (l *UiListener) process(s *api.Scan) {
	log.Debug("Scan received in UI listener", log.Any("scan", s))

	l.ui.AddScan(s)

	log.Debug("Scan displayed in UI listener", log.Any("scan", s))
}
