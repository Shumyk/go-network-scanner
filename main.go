package main

import (
	"go-network-scanner/api"
	"go-network-scanner/db"
	"go-network-scanner/pool"
	"go-network-scanner/ui"
	"go-network-scanner/util"

	"github.com/webui-dev/go-webui/v2"
	"go.uber.org/fx"
)

func main() {
	go util.SetupLogger()
	// 192.168.0.0/24
	databaseName, config := util.AskInput()

	app := fx.New(
		fx.Supply(&config),
		fx.Provide(
			db.NewDb(databaseName),
			db.NewConfigRepo,
			db.NewHostRepo,
			db.NewScanRepo,
			pool.NewHub,
			pool.NewWorkers,

			fx.Annotate(
				db.NewDbScanListener,
				fx.ResultTags(`group:"listener"`, `name:"db.scan"`),
			),
			fx.Annotate(
				db.NewDbHostListener,
				fx.ResultTags(`group:"listener"`, `name:"db.host"`),
			),
			fx.Annotate(
				ui.NewUiListener,
				fx.ResultTags(`group:"listener"`, `name:"ui"`),
			),
		),
		fx.Invoke(
			db.Bootstrap,
			fx.Annotate(
				api.SubscribeListeners,
				fx.ParamTags(`group:"listener"`),
			),
			pool.Start,
		),
	)

	app.Run()
	<-app.Done()

	w := webui.NewWindow()
	w.Show("<html><body><h1>Hello, folks!</h1>" + "</body></html>")
	webui.Wait()
}
