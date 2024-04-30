package util

import (
	"fmt"
	"go-network-scanner/api"
	"os"
	"strings"
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
	fmt.Println("Enter the second IP range to scan, optional: ")
	fmt.Scanln(&config.IpRange2)
	// 5. Third ip range to scan (optional)
	fmt.Println("Enter the third IP range to scan, optional: ")
	fmt.Scanln(&config.IpRange3)
	// 6. Enable socks proxy button and the ip address of the socks proxy, with a different proxy available for each ip range above
	askSocksProxyEnabled(&config)

	return
}

func askDatabaseName() string {
	fmt.Println("Enter the name of the sqlite database (defaults to 'scans'): ")
	var databaseName string
	fmt.Scanln(&databaseName)
	if len(databaseName) == 0 {
		return "test_db" // TODO: change to scans
	}
	return databaseName
}

func askSessionName(config *api.Config) {
	fmt.Println("Enter the name of the session: ")
	fmt.Scanln(&config.Name)
	if len(config.Name) == 0 {
		fmt.Fprintln(os.Stderr, "Session name is required")
		askSessionName(config)
	}
}

func askIpRange(config *api.Config) {
	fmt.Println("Enter the first IP range to scan: ")
	fmt.Scanln(&config.IpRange1)
	if len(config.IpRange1) == 0 {
		fmt.Fprintln(os.Stderr, "First IP range is required")
		askIpRange(config)
	}
}

func askSocksProxyEnabled(config *api.Config) {
	fmt.Println("Enable SOCKS proxy? (y/n): ")
	var enabled string
	fmt.Scanln(&enabled)

	enabled = strings.ToLower(strings.TrimSpace(enabled))
	config.EnableProxy = enabled[0] == 'y'
}
