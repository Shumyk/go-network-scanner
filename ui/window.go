// Package ui provides the user interface.
//
//nolint:wrapcheck // It's a wrapper for the webui.Window struct.
package ui

import "github.com/webui-dev/go-webui/v2"

type window interface {
	SetRootFolder(path string)
	ShowBrowser(content string, browser webui.Browser) error
	Show(content string) error
	Run(script string)
	Webui() webui.Window
}

type uiWindow struct {
	webui.Window
}

func (w *uiWindow) Webui() webui.Window {
	return w.Window
}
