package ui

import (
	"go-network-scanner/api"

	log "log/slog"

	"github.com/webui-dev/go-webui/v2"
)

const indexHTML = "index.html"

// NewUIListener creates a new UI listener.
func NewUIListener(ui *Ui) api.Listener {
	return &UiListener{ui}
}

// NewWebUI creates a new web UI and binds the functions to the buttons.
func NewWebUI() *Ui {
	log.Info("Creating UI")

	w := webui.NewWindow()
	ui := &Ui{&w}
	ui.window.SetRootFolder("ui/pages/build")

	webui.Bind(*ui.window, "bootstrap", ui.bootstrap)

	return ui
}

func Run(ui *Ui) {
	go ui.Run()
}
