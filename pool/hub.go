package pool

import (
	log "log/slog"
	"time"

	"github.com/sourcegraph/conc/pool"
)

var interval = 5 * time.Second

// Hub is the main struct that holds all workers.
type Hub struct {
	workers []Worker
}

// Start starts the ticker and begins cycle
// where it concurrently triggers all workers to do their scanning.
func Start(h *Hub) {
	go h.start()
}

func (h *Hub) start() {
	for ; true; <-time.NewTicker(interval).C {
		start := time.Now()

		log.Info("New pool cycle started")
		h.cycle()
		log.Info(
			"Pool cycle finished",
			log.Duration("duration", time.Since(start)),
		)
	}
}

func (h *Hub) cycle() {
	p := pool.New().WithMaxGoroutines(len(h.workers))
	for _, w := range h.workers {
		p.Go(w.Do)
	}
	p.Wait()
}
