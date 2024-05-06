package db

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
)

type HostListener struct {
	hosts *HostRepo
}

func (l *HostListener) Topic() events.EventName {
	return api.EVENT_HOST
}

func (l *HostListener) Listen(i ...interface{}) {
	go l.process(i[0].(api.Host))
}

func (l *HostListener) process(host api.Host) {
	log.Debug("Host received", log.Any("host", host))

	id, err := l.hosts.Upsert(host)
	if err != nil {
		log.Error("Failed to persist host in DB", log.Any("error", err))
		return
	}

	log.Debug(
		"Host persisted",
		log.Any("host.ip", host.Ip),
		log.Int("id", id),
	)
}
