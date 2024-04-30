package util

import (
	"fmt"
	"go-network-scanner/api"
	log "log/slog"
)

func AskInput() (databaseName string, config api.Config) {
	config = api.Config{}

	// 1. Name of session
	askSessionName(&config)
	// 2. Name of sqlite database (default to scans)
	databaseName = askDatabaseName()
	// 3. First ip range to scan (required)
	askIpRange(&config)
	// 4. Second ip range scan (optional)
	log.Info("Enter the second IP range to scan, optional: ")
	fmt.Scanln(&config.IpRange2)
	// 5. Third ip range to scan (optional)
	log.Info("Enter the third IP range to scan, optional: ")
	fmt.Scanln(&config.IpRange3)
	// 6. Enable socks proxy button and the ip address of the socks proxy, with a different proxy available for each ip range above
	log.Info("Enable SOCKS proxy? (y/n): ")
	fmt.Scanln(&config.EnableProxy)

	return
}

func askDatabaseName() string {
	log.Info("Enter the name of the sqlite database (defaults to 'scans'): ")
	var databaseName string
	fmt.Scanln(&databaseName)
	if len(databaseName) == 0 {
		return "scans"
	}
	return databaseName
}

func askSessionName(config *api.Config) {
	log.Info("Enter the name of the session: ")
	fmt.Scanln(&config.Name)
	if len(config.Name) == 0 {
		log.Error("Session name is required")
		askSessionName(config)
	}
}

func askIpRange(config *api.Config) {
	log.Info("Enter the first IP range to scan: ")
	fmt.Scanln(&config.IpRange1)
	if len(config.IpRange1) == 0 {
		log.Error("First IP range is required")
		askIpRange(config)
	}
}
