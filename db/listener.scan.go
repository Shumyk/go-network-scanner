package db

import (
	"go-network-scanner/api"
	log "log/slog"

	"github.com/kataras/go-events"
	"github.com/pkg/errors"
)

// ScanListener listens for scan events and persists them.
type ScanListener struct {
	scans RepoBasic[*api.Scan]
}

// Topic returns the event topic.
func (*ScanListener) Topic() events.EventName {
	return api.EventNmapScan
}

// Listen listens for scan events.
func (l *ScanListener) Listen(i ...any) {
	if len(i) == zero || i[zero] == nil {
		log.Error("no event data")
		return
	}
	s, ok := i[zero].(*api.Scan)
	if !ok {
		log.Error("Failed to cast to scan")
		return
	}
	//nolint:errcheck // Ignore error here as it's returned for tests.
	go l.process(s)
}

func (l *ScanListener) process(scan *api.Scan) (int, error) {
	log.Debug("Scan received", log.Any("scan", scan))

	id, err := l.scans.Upsert(scan)
	if err != nil {
		log.Error("failed to persist scan", log.Any("error", err))
		return id, errors.Wrap(err, "failed to persist scan")
	}

	log.Debug(
		"Scan persisted",
		log.Any("ip", scan.IP),
		log.Any("port", scan.Port),
		log.Int("id", id),
	)
	return id, nil
}
