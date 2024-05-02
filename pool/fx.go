package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	"slices"

	log "log/slog"
)

func NewHub(workers []*Worker) *Hub {
	return &Hub{workers}
}

func NewWorkers(c *api.Config, nmap *nmap.Scanner) []*Worker {
	ipRanges := c.IpRanges()
	log.Debug("Creating workers", log.Any("ipRanges", ipRanges))

	workers := make([]*Worker, len(ipRanges))
	for x, ip := range ipRanges {
		workers[x] = newWorker(ip, c.EnableProxy, nmap)
	}

	workers = slices.Clip(workers)
	log.Debug("Workers created", log.Any("workers", workers), log.Int("count", len(workers)))
	return workers
}

func newWorker(ipRange string, proxy bool, nmap *nmap.Scanner) *Worker {
	return &Worker{
		ipRange:     ipRange,
		enableSocks: proxy,
		nmap:        nmap,
	}
}
