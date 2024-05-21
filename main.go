// Package main : Main entry of the application
package main

import (
	"go-network-scanner/api"
	"go-network-scanner/db"
	"go-network-scanner/pool"
	"go-network-scanner/ui"
	"go-network-scanner/util"

	"go.uber.org/fx"
)

func main() {
	go util.SetupLogger()
	// 192.168.0.0/24

	u := ui.NewWebUI()
	go u.Run()
	conf := <-u.Conf

	app := fx.New(
		fx.Supply(
			api.ConfigOf(conf),
			u,
		),
		fx.Provide(
			db.NewDb(conf.DatabaseName),
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
				ui.NewUIListener,
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
}
