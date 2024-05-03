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
	return "scan"
}

func (l *HostListener) Listen(i ...interface{}) {
	go l.process(i[0].(api.ScanFound))
}

func (l *HostListener) process(scan api.ScanFound) {
	log.Info("Host received", log.Any("scan", scan))

	host := api.HostOf(scan)
	id, err := l.hosts.Upsert(host)
	if err != nil {
		log.Error("Failed to persist host in DB", log.Any("error", err))
		return
	}
	log.Info(
		"Host persisted",
		log.Any("host.ip", host.Ip),
		log.Int("id", id),
	)
}
