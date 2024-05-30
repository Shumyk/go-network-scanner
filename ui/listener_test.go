package ui

import (
	"go-network-scanner/api"
	"sync"
	"testing"

	"github.com/kataras/go-events"
	"github.com/stretchr/testify/assert"
)

type uiMock struct {
	wg     sync.WaitGroup
	called bool
}

func (u *uiMock) addScan(s *api.ScannedResultsDTO) {
	u.called = true
	u.wg.Done()
}

func TestTopic(t *testing.T) {
	l := &Listener{}
	assert.Equal(t, events.EventName(api.EventNmapScans), l.Topic())
}

func TestListen(t *testing.T) {
	tests := []struct {
		name   string
		data   any
		called bool
	}{
		{
			name:   "no data",
			data:   nil,
			called: false,
		},
		{
			name:   "invalid data",
			data:   "invalid",
			called: false,
		},
		{
			name:   "valid data",
			data:   &api.ScannedResultsDTO{},
			called: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &uiMock{}
			if tt.called {
				u.wg.Add(1)
			}

			l := &Listener{ui: u}
			l.Listen(tt.data)

			u.wg.Wait()
			assert.Equal(t, tt.called, u.called)
		})
	}
}
