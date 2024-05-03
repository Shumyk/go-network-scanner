package api

func HostOf(found ScanFound) Host {
	return Host{
		Ip:            found.Ip,
		OsType:        found.OsType,
		HostnameLabel: found.HostnameLabel,
		CustomLabel:   found.CustomLabel,
	}
}

func ScanOf(found ScanFound) Scan {
	return Scan{
		Ip:            found.Ip,
		Port:          found.Port,
		ServiceBanner: found.ServiceBanner,
	}
}
