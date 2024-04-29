package main

import (
	"log"

	"github.com/webui-dev/go-webui/v2"
)

func main() {
	setupLogger()

	w := webui.NewWindow()
	w.Show("<html><body><h1>Hello, folks!</h1></body></html>")
	webui.Wait()
}

func setupLogger() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
