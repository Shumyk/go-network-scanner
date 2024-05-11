package db

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
	"github.com/pkg/errors"
)

const zero = 0

// HostListener listens for host events and persists them.
type HostListener struct {
	hosts RepoBasic[*api.Host]
}

// Topic returns the event topic.
func (*HostListener) Topic() events.EventName {
	return api.EventNmapHost
}

// Listen listens for host events.
func (l *HostListener) Listen(i ...any) {
	if len(i) == zero || i[zero] == nil {
		log.Error("no event data")
		return
	}
	h, ok := i[zero].(*api.Host)
	if !ok {
		log.Error("failed to cast to host")
		return
	}
	go l.process(h)
}

func (l *HostListener) process(host *api.Host) (int, error) {
	log.Debug("host received", log.Any("host", host))

	id, err := l.hosts.Upsert(host)
	if err != nil {
		log.Error(
			"failed to persist host in DB",
			log.Any("error", err),
		)
		return id, errors.Wrap(err, "failed to persist host in DB")
	}

	log.Debug(
		"host persisted",
		log.Any("host.ip", host.IP),
		log.Int("id", id),
	)
	return id, nil
}
