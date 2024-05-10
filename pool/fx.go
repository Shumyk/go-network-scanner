package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	"go-network-scanner/util"

	log "log/slog"
)

func NewHub(workers []*Worker) *hub {
	return &hub{workers}
}

func NewWorkers(c *api.Config) []*Worker {
	ipRanges := c.IPRanges()
	log.Debug("Creating workers", log.Any("ipRanges", ipRanges))

	workers := make([]*Worker, len(ipRanges))
	for x, ip := range ipRanges {
		workers[x] = newWorker(ip, c.EnableProxy)
	}

	log.Debug("Workers created", log.Any("workers", workers), log.Int("count", len(workers)))
	return workers
}

func newWorker(ipRange string, proxy bool) *Worker {
	return &Worker{
		ipRange:     ipRange,
		enableSocks: proxy,
		nmap:        nmap.NewScanner(ipRange),
		log:         util.LoggerWith("worker", log.String("range", ipRange)),
	}
}
