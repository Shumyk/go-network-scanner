package ui

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

const (
	topic = "topic"
	data  = "data"
)

// Listener listens for events and propagates them to the UI.
type Listener struct {
	ui *UI
}

// Topic returns the event topic.
func (*Listener) Topic() events.EventName {
	return api.EventNmapScans
}

// Listen listens for events.
func (l *Listener) Listen(i ...any) {
	if len(i) == 0 || i[0] == nil {
		log.Error(
			"no event data",
			log.Any(topic, l.Topic()),
		)
		return
	}
	s, ok := i[0].(api.ScannedResultsDTO)
	if !ok {
		log.Error(
			"failed to cast",
			log.Any(topic, l.Topic()),
			log.Any(data, i[0]),
		)
		return
	}
	go l.ui.addScan(&s)
}
