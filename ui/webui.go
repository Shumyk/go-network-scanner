package ui

import (
	"encoding/json"
	"go-network-scanner/api"
	log "log/slog"
	"strconv"

	"github.com/webui-dev/go-webui/v2"
)

// UI represents the user interface.
type UI struct {
	window *webui.Window
	Conf   chan *api.ConfDTO
}

// 192.168.0.0/24

func (u *UI) Run() {
	log.Info("Running UI")
	if err := u.window.ShowBrowser(indexHTML, webui.ChromiumBased); err != nil {
		log.Warn("Install a chromium-based browser for an optimized experience", log.Any("err", err))
		err := u.window.Show(indexHTML)
		if err != nil {
			log.Error("Failed to show window", log.Any("err", err))
		}
	}
	webui.Wait()
}

func (u *UI) addScan(s map[string][]*api.ScanDTO) {
	log.Info(
		"Adding scans to UI",
		log.Any("scan.ip", len(s)),
	)

	bytes, err := json.Marshal(s)
	if err != nil {
		log.Error("Failed to marshal scans", log.Any("err", err))
		return
	}

	u.window.Run("addScannedDevice('" + string(bytes) + "')")
}

func countFunc(e webui.Event) any {
	log.Info("Button 1 clicked")
	count, _ := e.Window.Script("return count;", webui.ScriptOptions{})
	i, _ := strconv.Atoi(count)
	e.Window.Run("SetCount(" + strconv.Itoa(i+1) + ")")
	return nil
}

func exitFunc(e webui.Event) any {
	log.Info("Button 2 clicked")
	webui.Exit()
	return nil
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
	}
	log.Info("Got config", log.Any("conf", conf))
	u.Conf <- &conf

	return nil
}
