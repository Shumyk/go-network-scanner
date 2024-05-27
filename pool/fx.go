// Package pool contains the worker hub and the worker struct.
package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	"go-network-scanner/util"

	log "log/slog"
)

// NewHub creates a new hub with the given workers.
func NewHub(workers []Worker) *Hub {
	return &Hub{workers}
}

// NewWorkers creates a new worker for each IP range in the config.
func NewWorkers(c *api.Config) []Worker {
	ipRanges := c.IPRanges()
	log.Debug("Creating workers", log.Any("ipRanges", ipRanges))

	workers := make([]Worker, len(ipRanges))
	for x, el := range ipRanges {
		workers[x] = newWorker(el.IPRange, el.Proxy)
	}

	log.Debug(
		"Workers created",
		log.Any("workers", workers),
		log.Int("count", len(workers)),
	)
	return workers
}

func newWorker(ipRange, proxy string) Worker {
	return &worker{
		nmap: nmap.NewScanner(ipRange, proxy),
		log:  util.LoggerWith("worker", log.String("range", ipRange)),
	}
}
