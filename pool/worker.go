package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	log "log/slog"

	"github.com/kataras/go-events"
)

type Worker struct {
	ipRange     string
	enableSocks bool
	nmap        *nmap.Scanner
}

func (w *Worker) Do() {
	log.Info("Starting scan", log.String("ipRange", w.ipRange))
	scanResult, err := w.nmap.Scan(w.ipRange)
	log.Info("Finished scan", log.String("ipRange", w.ipRange))

	if err != nil {
		log.Error("Error scanning", log.String("ipRange", w.ipRange), log.Any("err", err))
		panic(err)
	}

	for _, h := range scanResult.Hosts {
		for _, p := range h.Ports {
			log.Info("Has addresses", log.Any("addresses", h.Addresses))
			log.Info("Has hostnames", log.Any("hostnames", h.Hostnames))

			hostname := ""
			if len(h.Hostnames) > 0 {
				hostname = h.Hostnames[0].Name
			}
			customLabel := ""
			if len(h.Hostnames) > 1 {
				log.Warn("More than one hostname")
				customLabel = h.Hostnames[1].Name
			}
			scan := api.ScanFound{
				Ip:            h.Addresses[0].Addr,
				Port:          p.ID,
				ServiceBanner: p.Service.Name,
				OsType:        p.Service.OSType,
				HostnameLabel: hostname,
				CustomLabel:   customLabel,
			}
			events.Emit("scan", scan)
		}
	}
}
