package pool

import (
	log "log/slog"
	"time"

	"github.com/sourcegraph/conc/pool"
)

type hub struct {
	workers []*Worker
}

// Start starts the ticker and begins cycle
// where it concurrently triggers all workers to do their scanning.
func Start(h *hub) {
	go h.start()
}

func (h *hub) start() {
	// TODO: take time from config
	for range time.NewTicker(5 * time.Second).C {
		start := time.Now()

		log.Info("New pool cycle started")
		h.cycle()
		log.Info(
			"Pool cycle finished",
			log.Duration("duration", time.Since(start)),
		)
	}
}

func (h *hub) cycle() {
	p := pool.New().WithMaxGoroutines(len(h.workers))
	for _, w := range h.workers {
		p.Go(w.Do)
	}
	p.Wait()
}
