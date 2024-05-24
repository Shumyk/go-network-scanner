package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
	log "log/slog"
	"time"

	lib "github.com/Ullaakut/nmap/v3"
	"github.com/kataras/go-events"
)

// Worker is the interface that all workers must implement.
type Worker interface {
	Do()
}

// worker is the main struct that does the scanning and emits events.
type worker struct {
	nmap nmap.Scanner
	log  *log.Logger

	enableSocks bool
}

// Do triggers the nmap scan and processes the results.
func (w *worker) Do() {
	scanResult := w.nmap.Scan()
	scannedAt := time.Now()

	dto := api.ScannedResultsDTO{
		Data:      []api.HostDTO{},
		ScannedAt: scannedAt,
	}

	for i := range scanResult.Hosts {
		h := &scanResult.Hosts[i]
		w.log.Debug("Host found", log.Any("host", h))

		ip := nmap.IP(h)
		host := processHost(ip, h)
		hostDTO := api.HostDTOFrom(host, scannedAt)

		for x := range h.Ports {
			scan := processPort(ip, &h.Ports[x])
			scanDTO := api.ScanDTOFrom(scan, scannedAt)
			hostDTO.Services = append(hostDTO.Services, *scanDTO)
		}
		dto.Data = append(dto.Data, *hostDTO)
	}
	events.Emit(api.EventNmapScans, dto)
}

func processHost(ip string, h *lib.Host) *api.Host {
	// nolint:exhaustruct // Other fields are added by consumer.
	host := api.Host{
		FirstDiscovered: time.Time(h.StartTime),
		LastSeen:        time.Time(h.EndTime),
		IP:              ip,
		OsType:          nmap.OS(h),
		HostnameLabel:   nmap.Hostname(h),
		CustomLabel:     h.Comment,
	}
	events.Emit(api.EventNmapHost, &host)
	return &host
}

func processPort(ip string, p *lib.Port) *api.Scan {
	// nolint:exhaustruct // Other fields are added by consumer.
	scan := api.Scan{
		FirstDiscovered: time.Now(),
		LastSeen:        time.Now(),
		IP:              ip,
		Port:            p.ID,
		ServiceBanner:   nmap.ServiceBanner(p),
	}
	events.Emit(api.EventNmapScan, &scan)
	return &scan
}
