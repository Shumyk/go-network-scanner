package api

import (
	log "log/slog"
	"reflect"

	"github.com/kataras/go-events"
)

const (
	// EventBootstrap is topic name for singular bootstrap event.
	EventBootstrap = "bootstrap"
	// EventNmapScan is topic name for scans (services running on port).
	EventNmapScan = "nmap.scan"
	// EventNmapHost is topic name for a hosts (IP address).
	EventNmapHost = "nmap.host"
)

// Listener is an interface for listening to events.
type Listener interface {
	Topic() events.EventName
	Listen(i ...any)
}

// SubscribeListeners subscribes the given listeners to their respective topics.
// It logs the subscription details and registers the listeners with
// the events package.
func SubscribeListeners(ll []Listener) {
	for _, l := range ll {
		log.Info(
			"Subscribing to",
			log.Any("topic", l.Topic()),
			log.Any("listener", reflect.TypeOf(l).Elem().Name()),
		)
		events.On(l.Topic(), l.Listen)
	}
}
