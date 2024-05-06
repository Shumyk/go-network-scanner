package pool

import (
	"fmt"
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	"log/slog"
	log "log/slog"

	"github.com/kataras/go-events"
)

type Worker struct {
	ipRange     string
	enableSocks bool

	nmap *nmap.Scanner
	log  *slog.Logger
}

func (w *Worker) Do() {
	w.log.Info("Starting scan")
	scanResult, err := w.nmap.Scan(w.ipRange)
	if err != nil {
		w.log.Error("Error scanning", log.Any("err", err))
		panic(err)
	}
	w.log.Info("Finished scan")

	for _, h := range scanResult.Hosts {
		w.log.Info("Host found", log.Any("host", h))

		hostname := ""
		if len(h.Hostnames) > 0 {
			hostname = h.Hostnames[0].Name
		}
		host := api.Host{
			Ip:            h.Addresses[0].Addr,
			OsType:        h.OS.Matches[0].Name,
			HostnameLabel: hostname,
			CustomLabel:   h.Comment,
		}
		events.Emit(api.EVENT_HOST, host)
		for _, p := range h.Ports {
			banner := fmt.Sprintf("%v - %v - %v", p.Service.Name, p.Service.Product, p.Service.Version)
			scan := api.Scan{
				Ip:            h.Addresses[0].Addr,
				Port:          p.ID,
				ServiceBanner: banner,
			}
			events.Emit(api.EVENT_SCAN, scan)
		}
	}
}
