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

type ConfDTO struct {
	DatabaseName string
	SessionName  string
	IPRange1     string
	IPRange2     string
	IPRange3     string
	EnableProxy  bool
	Proxy1       string
	Proxy2       string
	Proxy3       string
}

// 192.168.0.0/24

func (u *Ui) Run() {
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

func (u *Ui) AddScan(s *api.Scan) {
	log.Info("Adding scan to UI", log.Any("scan.ip", s.IP))
	u.window.Run("addScannedDevice('" + s.IP + ", " + strconv.Itoa(int(s.Port)) + ", " + s.ServiceBanner + "')")
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

func (*Ui) bootstrap(e webui.Event) webui.Void {
	log.Info("Bootstrapping UI",
		log.Any("event", e),
		log.Any("event.element", e.Element),
		log.Any("event.elementType", e.EventType),
	)

	conf, err := webui.GetArg[ConfDTO](e)
	if err != nil {
		log.Error("Failed to get config for bootstrap", log.Any("err", err))
	}
	log.Info("Got config", log.Any("conf", conf))

	return nil
}
