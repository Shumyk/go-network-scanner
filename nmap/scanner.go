package nmap

import (
	"context"
	log "log/slog"

	"github.com/Ullaakut/nmap/v3"
)

func Scan(ip string) {
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(ip),
		nmap.WithServiceInfo(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
	)
	if err != nil {
		log.Error("Error creating scanner: %v", err, log.String("ip", ip))
		return
	}

	results, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Warn("Warnings while scanning: %v", warnings)
	}
	if err != nil {
		log.Error("Error scanning: %v", err)
		return
	}

	for _, host := range results.Hosts {
		log.Info("Host %s (%s)", host.Addresses[0], host.Status.State)
		for _, port := range host.Ports {
			log.Info("\tPort %d/%s %s %s", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}
