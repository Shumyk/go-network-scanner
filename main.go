package main

import (
	"go-network-scanner/db"
	"go-network-scanner/nmap"
	"go-network-scanner/pool"
	"go-network-scanner/util"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/webui-dev/go-webui/v2"
	"go.uber.org/fx"
)

func main() {
	setupLogger()
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

func setupLogger() {
	l := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Key == slog.SourceKey {
						source, _ := a.Value.Any().(*slog.Source)
						if source != nil {
							// directory / filename.go
							source.File = filepath.Join(
								filepath.Base(filepath.Dir(source.File)),
								filepath.Base(source.File),
							)
						}
					}
					return a
				},
			},
		),
	)
	slog.SetDefault(l)
}
