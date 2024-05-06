package nmap

import (
	"context"
	log "log/slog"

	"github.com/Ullaakut/nmap/v3"
)

func NewScanner(ipRange string) *Scanner {
	return &Scanner{
		ip:      ipRange,
		scanner: newNmapScanner(ipRange),
	}
}

func newNmapScanner(ipRange string) *nmap.Scanner {
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(ipRange),
		nmap.WithServiceInfo(),
		nmap.WithOSDetection(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
	)
	if err != nil {
		log.Error(
			"Error creating scanner",
			log.Any("err", err),
			log.String("ip", ipRange),
		)
		panic(err)
	}
	return scanner
}
