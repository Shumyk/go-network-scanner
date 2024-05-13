package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	log "log/slog"

	nmaplib "github.com/Ullaakut/nmap/v3"
	"github.com/kataras/go-events"
)

type Worker struct {
	ipRange     string // TODO: is needed?
	enableSocks bool

	nmap *nmap.Scanner
	log  *log.Logger
}

func (w *Worker) Do() {
	w.log.Info("Starting scan")
	scanResult := w.nmap.Scan()
	w.log.Info("Finished scan")

	for _, h := range scanResult.Hosts {
		w.log.Info("Host found", log.Any("host", h))

		ip := nmap.Ip(h)
		processHost(ip, h)
		for _, p := range h.Ports {
			processPort(ip, p)
		}
	}
}

func processHost(ip string, h nmaplib.Host) {
	host := api.Host{
		IP:            ip,
		OsType:        nmap.Os(h),
		HostnameLabel: nmap.Hostname(h),
		CustomLabel:   h.Comment,
	}
	events.Emit(api.EventNmapHost, &host)
}

func processPort(ip string, p nmaplib.Port) {
	scan := api.Scan{
		IP:            ip,
		Port:          p.ID,
		ServiceBanner: nmap.ServiceBanner(p),
	}
	events.Emit(api.EventNmapScan, &scan)
}
