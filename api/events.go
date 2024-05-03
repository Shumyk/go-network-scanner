package api

import (
	log "log/slog"

	"github.com/kataras/go-events"
)

type Listener interface {
	Topic() events.EventName
	Listen(i ...interface{})
}

func SubscribeListeners(listeners []Listener) {
	for _, l := range listeners {
		log.Info("Subscribing to", log.Any("topic", l.Topic()), log.Any("listener", l))
		events.On(l.Topic(), l.Listen)
	}
}
