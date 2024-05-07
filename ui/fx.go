package ui

import (
	"go-network-scanner/api"

	log "log/slog"

	"github.com/webui-dev/go-webui/v2"
)

const indexHTML = "index.html"

func NewUiListener(ui *Ui) api.Listener {
	return &UiListener{ui}
}

func NewWebUi() *Ui {
	log.Info("Creating UI")

	w := webui.NewWindow()
	w.SetRootFolder("ui/pages")
	w.Bind("MyButton1", countFunc)
	w.Bind("MyButton2", exitFunc)

	return &Ui{&w}
}

func Run(ui *Ui) {
	go ui.Run()
}
