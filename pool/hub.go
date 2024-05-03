package pool

import (
	log "log/slog"
	"time"

	"github.com/sourcegraph/conc/pool"
)

type Hub struct {
	workers []*Worker
}

func Start(h *Hub) {
	h.start()
}

func (h *Hub) start() {
	// TODO: take time from config
	for range time.NewTicker(5 * time.Second).C {
		log.Info("New pool cycle started")
		start := time.Now()
		pool := pool.New().WithMaxGoroutines(len(h.workers))
		poolCreationDur := time.Since(start)

		for _, w := range h.workers {
			pool.Go(w.Do)
		}
		pool.Wait()

		log.Info(
			"Pool cycle finished",
			log.Duration("poolCreation", poolCreationDur),
			log.Duration("poolCycle", time.Since(start)),
		)
	}
}
