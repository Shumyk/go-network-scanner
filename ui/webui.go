package ui

import (
	"go-network-scanner/api"
	log "log/slog"
	"strconv"

	"github.com/webui-dev/go-webui/v2"
)

type Ui struct {
	window *webui.Window
}

func (ui *Ui) Run() {
	log.Info("Running UI")
	if err := ui.window.ShowBrowser(indexHTML, webui.ChromiumBased); err != nil {
		log.Warn("Install a chromium-based browser for an optimized experience")
		ui.window.Show(indexHTML)
	}
	webui.Wait()
}

func (u *Ui) AddScan(s api.Scan) {
	log.Info("Adding scan to UI", log.Any("scan.ip", s.IP))
	u.window.Run("AddDevice('" + s.IP + ", " + strconv.Itoa(int(s.Port)) + ", " + s.ServiceBanner + "')")
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
