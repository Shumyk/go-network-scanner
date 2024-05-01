package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	"slices"

	"github.com/sourcegraph/conc/pool"
)

func NewHub(c *api.Config, workers []*Worker) *Hub {
	return &Hub{
		c,
		pool.
			NewWithResults[api.Scan]().
			WithMaxGoroutines(c.IpRangesCount()),
		workers,
	}
}

func NewWorkers(c *api.Config, nmap *nmap.Scanner) []*Worker {
	ipRanges := c.IpRanges()
	workers := make([]*Worker, len(ipRanges))
	for _, ip := range ipRanges {
		w := newWorker(ip, c.EnableProxy, nmap)
		workers = append(workers, w)
	}
	return slices.Clip(workers)
}

func newWorker(ipRange string, proxy bool, nmap *nmap.Scanner) *Worker {
	return &Worker{
		ipRange:     ipRange,
		enableSocks: proxy,
		nmap:        nmap,
	}
}
