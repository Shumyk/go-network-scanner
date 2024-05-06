package api

import (
	log "log/slog"
	"reflect"

	"github.com/kataras/go-events"
)

const (
	EVENT_SCAN = "nmap.scan"
	EVENT_HOST = "nmap.host"
)

type Listener interface {
	Topic() events.EventName
	Listen(i ...interface{})
}

// SubscribeListeners subscribes the given listeners to their respective topics.
// It logs the subscription details and registers the listeners with the events package.
func SubscribeListeners(listeners []Listener) {
	for _, l := range listeners {
		log.Info(
			"Subscribing to",
			log.Any("topic", l.Topic()),
			log.Any("listener", reflect.TypeOf(l).Elem().Name()),
		)
		events.On(l.Topic(), l.Listen)
	}
}
