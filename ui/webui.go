package ui

import (
	"encoding/json"
	"go-network-scanner/api"
	log "log/slog"

	"github.com/webui-dev/go-webui/v2"
)

type uiListener interface {
	addScan(s *api.ScannedResultsDTO)
}

// UI represents the user interface.
type UI struct {
	window window
	Conf   chan *api.ConfDTO
}

// Run starts the user interface.
func (u *UI) Run() {
	log.Info("Running UI")
	if err := u.window.ShowBrowser(indexHTML, webui.ChromiumBased); err != nil {
		log.Warn(
			"Install a chromium-based browser for an optimized experience",
			log.Any("err", err),
		)
		showErr := u.window.Show(indexHTML)
		if showErr != nil {
			log.Error("Failed to show window", log.Any("err", err))
			panic(showErr)
		}
	}
	webui.Wait()
}

func (u *UI) addScan(s *api.ScannedResultsDTO) {
	log.Info(
		"Adding scans to UI",
		log.Any("count", len(s.Data)),
	)

	bytes, err := json.Marshal(s)
	if err != nil {
		log.Error("Failed to marshal scans", log.Any("err", err))
		return
	}

	u.window.Run("addScannedDevice('" + string(bytes) + "')")
}

func (u *UI) bootstrap(e webui.Event) webui.Void {
	log.Info("Bootstrapping UI",
		log.Any("event", e),
		log.Any("event.element", e.Element),
		log.Any("event.elementType", e.EventType),
	)

	conf, err := webui.GetArg[api.ConfDTO](e)
	if err != nil {
		log.Error("Failed to get config for bootstrap", log.Any("err", err))
		panic(err)
	}
	log.Info("Got config", log.Any("conf", conf))
	u.Conf <- &conf

	return nil
}
