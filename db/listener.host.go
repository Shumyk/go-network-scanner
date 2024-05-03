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

func (l *HostListener) process(found api.ScanFound) {
	// Persist the host in the database
	host := api.Host{
		Ip:            found.Ip,
		OsType:        found.OsType,
		HostnameLabel: found.HostnameLabel,
		CustomLabel:   found.CustomLabel,
	}
	id, err := l.hosts.Insert(host)
	if err != nil {
		log.Error("Failed to persist host in DB", log.Any("error", err))
		return
	}
	log.Info(
		"Host persisted in DB listener",
		log.Any("host.ip", host.Ip),
		log.Int("id", id),
	)
}
