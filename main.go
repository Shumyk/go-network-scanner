package main

import (
	"go-network-scanner/db"
	"go-network-scanner/nmap"
	"go-network-scanner/util"

	"github.com/webui-dev/go-webui/v2"
	"go.uber.org/fx"
)

func main() {
	// nmap.Scan("192.168.0.0/24")
	databaseName, config := util.AskInput()

	app := fx.New(
		fx.Provide(
			fx.Supply(&config),
			db.NewDb(databaseName),
			db.NewConfigRepo,
		),
		fx.Invoke(db.Bootstrap),
	)

	go nmap.Scan(config.IpRange1)
	if len(config.IpRange2) > 0 {
		go nmap.Scan(config.IpRange2)
	}
	if len(config.IpRange3) > 0 {
		go nmap.Scan(config.IpRange3)
	}

	app.Run()
	<-app.Done()

	w := webui.NewWindow()
	w.Show("<html><body><h1>Hello, folks!</h1>" + "</body></html>")
	webui.Wait()
}
