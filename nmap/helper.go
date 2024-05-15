package nmap

import "github.com/Ullaakut/nmap/v3"

const separator = " - "

// IP returns the first IP address of a host.
func IP(h *nmap.Host) string {
	if len(h.Addresses) > 0 {
		return h.Addresses[0].Addr
	}
	return ""
}

// Hostname returns the first hostname of a host.
func Hostname(h *nmap.Host) string {
	if len(h.Hostnames) > 0 {
		return h.Hostnames[0].Name
	}
	return ""
}

// OS returns the first OS name of a host.
func OS(h *nmap.Host) string {
	if len(h.OS.Matches) > 0 {
		return h.OS.Matches[0].Name
	}
	return ""
}

// ServiceBanner returns a string representation of a service.
func ServiceBanner(p *nmap.Port) string {
	banner := ""
	if p.Service.Name != "" {
		banner += p.Service.Name
	}
	if p.Service.Product != "" {
		banner = separateIfNotEmpty(banner) + p.Service.Product
	}
	if p.Service.Version != "" {
		banner = separateIfNotEmpty(banner) + p.Service.Version
	}
	return banner
}

func separateIfNotEmpty(s string) string {
	if s != "" {
		return s + separator
	}
	return s
}
