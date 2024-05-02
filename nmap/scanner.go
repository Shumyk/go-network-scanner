package nmap

import (
	"context"
	log "log/slog"

	"github.com/Ullaakut/nmap/v3"
)

type Scanner struct{}

// nmap -Av 192.168.0.0/24
func (s *Scanner) Scan(ip string) (*nmap.Run, error) {
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(ip),
		nmap.WithServiceInfo(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
	)
	if err != nil {
		log.Error("Error creating scanner: %v", err, log.String("ip", ip))
		return nil, err
	}

	log.Info("Nmap starts scanning", log.String("ip", ip))
	results, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Warn("Warnings while scanning: %v", warnings)
	}
	if err != nil {
		log.Error("Error scanning: %v", log.String("ip", ip), log.Any("err", err))
		return nil, err
	}

	for _, host := range results.Hosts {
		os := "Unknown"
		hostname := "Unknown"
		if len(host.OS.Matches) > 0 {
			os = host.OS.Matches[0].Name
		}
		if len(host.Hostnames) > 0 {
			hostname = host.Hostnames[0].Name
		}
		log.Info(
			"Host: ",
			log.String("ip", host.Addresses[0].Addr),
			log.String("state", host.Status.State),
			log.String("os", os),
			log.String("hostname", hostname),
		)
		for _, port := range host.Ports {
			log.Info(
				"\t\tPort: ",
				log.Uint64("id", uint64(port.ID)),
				log.String("protocol", port.Protocol),
				log.String("state", port.State.State),
				log.String("name", port.Service.Name),
				log.String("product", port.Service.Product),
				log.String("version", port.Service.Version),
				log.String("extrainfo", port.Service.ExtraInfo),
				log.String("reason", port.State.Reason),
				log.String("os", port.Service.OSType),
			)
		}
	}
	return results, nil
}
