//nolint:forbidigo,revive // This file is used to ask for user input, so it's okay to ignore warning about fmt.* calls.
package util

import (
	"fmt"
	"go-network-scanner/api"
	"os"
	"strings"
)

const (
	DEFAULT_TEST_DB = "test_db" // TODO: change to scans
)

var input = os.Stdin

func AskInput() (databaseName string, config api.Config) {
	config = api.EmptyConfig()

	// 1. Name of session
	askSessionName(&config)

	// 2. Name of sqlite database (default to scans)
	databaseName = askDatabaseName()

	// 3. First ip range to scan (required)
	askIpRange(&config)

	// 4. Second ip range scan (optional)
	fmt.Println("Enter the second IP range to scan, optional: ")
	fmt.Fscanln(input, &config.IPRange2)

	// 5. Third ip range to scan (optional)
	fmt.Println("Enter the third IP range to scan, optional: ")
	fmt.Fscanln(input, &config.IPRange3)

	// 6. Enable socks proxy button and the ip address of the socks proxy, with a different proxy available for each ip range above
	askSocksProxyEnabled(&config)

	return
}

func askDatabaseName() string {
	fmt.Println("Enter the name of the sqlite database (defaults to 'scans'): ")
	var databaseName string
	fmt.Fscanln(input, &databaseName)

	if databaseName == "" {
		return DEFAULT_TEST_DB
	}
	return strings.TrimSpace(databaseName)
}

func askSessionName(config *api.Config) {
	fmt.Println("Enter the name of the session: ")
	fmt.Fscanln(input, &config.Name)

	if config.Name == "" {
		fmt.Fprintln(os.Stderr, "Session name is required")
		askSessionName(config)
	}
}

func askIpRange(config *api.Config) {
	fmt.Println("Enter the first IP range to scan: ")
	fmt.Fscanln(input, &config.IPRange1)

	if config.IPRange1 == "" {
		fmt.Fprintln(os.Stderr, "First IP range is required")
		askIpRange(config)
	}
}

func askSocksProxyEnabled(config *api.Config) {
	fmt.Println("Enable SOCKS proxy? [y/n]: ")
	var enabled string
	fmt.Fscanln(input, &enabled)

	enabled = strings.ToLower(strings.TrimSpace(enabled))
	config.EnableProxy = enabled != "" && enabled[0] == 'y'
}
