package nmap

import "github.com/Ullaakut/nmap/v3"

const SEPARATOR = " - "

func Ip(h nmap.Host) string {
	if len(h.Addresses) > 0 {
		return h.Addresses[0].Addr
	}
	return ""
}

func Hostname(h nmap.Host) string {
	if len(h.Hostnames) > 0 {
		return h.Hostnames[0].Name
	}
	return ""
}

func Os(h nmap.Host) string {
	if len(h.OS.Matches) > 0 {
		return h.OS.Matches[0].Name
	}
	return ""
}

func ServiceBanner(p nmap.Port) string {
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
		return s + SEPARATOR
	}
	return s
}
