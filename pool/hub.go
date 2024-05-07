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
	go h.start()
}

func (h *Hub) start() {
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

func (h *Hub) cycle() {
	pool := pool.New().WithMaxGoroutines(len(h.workers))
	for _, w := range h.workers {
		pool.Go(w.Do)
	}
	pool.Wait()
}
