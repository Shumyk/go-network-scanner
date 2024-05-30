package ui

import (
	"encoding/json"
	"errors"
	"go-network-scanner/api"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webui-dev/go-webui/v2"
)

type webuiMock struct {
	runArg         string
	showBrowserErr error
	showErr        error
}

func (w *webuiMock) Run(script string) {
	w.runArg = script
}
func (w *webuiMock) SetRootFolder(path string) {}
func (w *webuiMock) ShowBrowser(content string, browser webui.Browser) error {
	return w.showBrowserErr
}
func (w *webuiMock) Show(content string) error {
	return w.showErr
}
func (w *webuiMock) Webui() webui.Window {
	return webui.NewWindow()
}

func TestRunFail(t *testing.T) {
	// Given
	mock := &webuiMock{
		showBrowserErr: errors.New("no chromium"),
		showErr:        errors.New("no window"),
	}
	ui := &UI{window: mock}
	// When / Then
	assert.PanicsWithError(t, "no window", ui.Run)
}

func TestRun(t *testing.T) {
	// Given
	mock := &webuiMock{
		showBrowserErr: errors.New("no chromium"),
	}
	ui := &UI{window: mock}
	// When / Then
	assert.NotPanics(t, ui.Run)
}

func TestAddScan(t *testing.T) {
	// Given
	mock := &webuiMock{}
	ui := &UI{window: mock}
	s := &api.ScannedResultsDTO{
		Data: []api.HostDTO{
			{
				IP:            "host-ip",
				OS:            "linux",
				HostnameLabel: "host.io",
				CustomLabel:   "custom label",
				Services: []api.ScanDTO{
					{
						Port:   80,
						IP:     "service-ip",
						Banner: "banner",
					},
				},
			},
		},
	}
	// When
	ui.addScan(s)
	// Then
	assert.Equal(
		t,
		`addScannedDevice('{"data":[{"firstDiscovered":"0001-01-01 00:00:00","lastDiscovered":"0001-01-01 00:00:00","ip":"host-ip","os":"linux","hostnameLabel":"host.io","customLabel":"custom label","services":[{"firstDiscovered":"0001-01-01 00:00:00","lastDiscovered":"0001-01-01 00:00:00","ip":"service-ip","banner":"banner","port":80}]}]}')`,
		mock.runArg,
	)
}

func TestBootstrapFail(t *testing.T) {
	// Given
	mock := &webuiMock{}
	confChan := make(chan *api.ConfDTO)
	ui := &UI{window: mock, Conf: confChan}

	input := api.ConfDTO{
		DatabaseName: "db",
		SessionName:  "session",
		IPRange1:     "ip1",
	}
	j, err := json.Marshal(input)
	assert.NoError(t, err)
	event := webui.Event{Element: string(j)}

	// When // Then
	assert.Panics(t, func() { ui.bootstrap(event) })
}
