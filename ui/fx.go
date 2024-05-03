package ui

import "go-network-scanner/api"

func NewUiListener() api.Listener {
	return &UiListener{}
}
