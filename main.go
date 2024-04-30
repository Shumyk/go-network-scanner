package main

import (
	"go-network-scanner/db"
	"go-network-scanner/nmap"
	"go-network-scanner/util"

	"github.com/webui-dev/go-webui/v2"
)

func main() {
	databaseName, config := util.AskInput()
	db.Init(databaseName, &config)

	nmap.Scan(config.IpRange1)
	if len(config.IpRange2) > 0 {
		nmap.Scan(config.IpRange2)
	}
	if len(config.IpRange3) > 0 {
		nmap.Scan(config.IpRange3)
	}

	w := webui.NewWindow()
	w.Show("<html><body><h1>Hello, folks!</h1>" + "</body></html>")
	webui.Wait()
}
