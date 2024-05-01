package pool

import (
	"go-network-scanner/api"
	"log/slog"
	"time"

	"github.com/sourcegraph/conc/pool"
)

type Hub struct {
	conf    *api.Config // TODO: is this needed?
	pool    *pool.ResultPool[api.Scan]
	workers []*Worker
}

func Start(h *Hub) {
	h.start()
}

func (h *Hub) start() {
	ticker := time.NewTicker(5 * time.Second) // TODO: take time from config
	// defer ticker.Stop()

	for range ticker.C {
		for _, w := range h.workers {
			h.pool.Go(w.Do)
		}
		scans := h.pool.Wait()
		slog.Info("Retrieved scans", slog.Any("scans", scans), slog.Int("count", len(scans)))
		// TODO: save scans to db
	}
}
