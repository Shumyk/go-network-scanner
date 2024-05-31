// Package ui provides the user interface.
package ui

import (
	"go-network-scanner/api"

	log "log/slog"

	"github.com/webui-dev/go-webui/v2"
)

// NewUIListener creates a new UI listener.
func NewUIListener(ui *UI) api.Listener {
	return &Listener{ui}
}

// NewWebUI creates a new web UI and binds the functions to the buttons.
func NewWebUI() *UI {
	log.Info("Creating UI")

	w := uiWindow{webui.NewWindow()}
	ui := &UI{&w, make(chan *api.ConfDTO)}

	webui.Bind(ui.window.Webui(), "bootstrap", ui.bootstrap)

	return ui
}
