package ui

import (
	"go-network-scanner/api"
	"strconv"

	log "log/slog"

	"github.com/webui-dev/go-webui/v2"
)

const indexHTML = "index.html"

func NewUiListener() api.Listener {
	return &UiListener{}
}

func NewWebUi() *webui.Window {
	log.Info("Creating UI")

	w := webui.NewWindow()
	w.SetRootFolder("ui/pages")
	w.Bind("MyButton1", countFunc)
	w.Bind("MyButton2", exitFunc)

	return &w
}

func Run(w *webui.Window) {
	log.Info("Running UI")
	if err := w.ShowBrowser(indexHTML, webui.ChromiumBased); err != nil {
		log.Warn("Install a chromium-based browser for an optimized experience")
		w.Show(indexHTML)
	}
	webui.Wait()
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
