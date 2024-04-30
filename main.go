package main

import (
	"go-network-scanner/db"
	"go-network-scanner/util"
	"log/slog"

	"github.com/webui-dev/go-webui/v2"
)

func main() {
	databaseName, config := util.AskInput()
	db.Init(databaseName, &config)
	configs, err := db.Configs.All()
	if nil != err {
		panic(err)
	}
	slog.Info("Configs: ", configs)

	configsStr := "["
	for _, config := range configs {
		configsStr += config.Name + ", " + config.IpRange1 + ", " + config.IpRange2 + ", " + config.IpRange3 + ", " + config.CreatedAt.String() + "\n"
	}
	configsStr += "]"
	slog.Info("Configs: ", configsStr)

	w := webui.NewWindow()
	w.Show("<html><body><h1>Hello, folks!</h1>" + configsStr + "</body></html>")
	webui.Wait()
}
