package main

import (
	"go-network-scanner/db"
	"go-network-scanner/nmap"
	"go-network-scanner/pool"
	"go-network-scanner/util"

	"github.com/webui-dev/go-webui/v2"
	"go.uber.org/fx"
)

func main() {
	// nmap.Scan("192.168.0.0/24")
	databaseName, config := util.AskInput()

	app := fx.New(
		fx.Supply(&config),
		fx.Provide(
			db.NewDb(databaseName),
			db.NewConfigRepo,
			nmap.NewScanner,
			pool.NewHub,
			pool.NewWorkers,
		),
		fx.Invoke(
			db.Bootstrap,
			pool.Start,
		),
	)

	app.Run()
	<-app.Done()

	w := webui.NewWindow()
	w.Show("<html><body><h1>Hello, folks!</h1>" + "</body></html>")
	webui.Wait()
}
