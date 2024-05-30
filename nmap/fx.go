// Package nmap is a wrapper around the nmap package.
// It provides a simple interface to scan a range of IP addresses.
package nmap

import (
	"context"
	log "log/slog"

	"github.com/Ullaakut/nmap/v3"
)

// NewScanner creates a new Scanner.
func NewScanner(ipRange, proxy string) Scanner {
	return &scanner{
		ip:      ipRange,
		scanner: newNmapScanner(ipRange, proxy),
	}
}

func newNmapScanner(ipRange, proxy string) scanRunner {
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(ipRange),
		nmap.WithServiceInfo(),
		nmap.WithOSDetection(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
	)
	if proxy != "" {
		scanner.AddOptions(nmap.WithProxies("socks4://" + proxy))
	}
	if err != nil {
		log.Error("Error creating scanner", log.Any("err", err))
		panic(err)
	}
	return scanner
}
