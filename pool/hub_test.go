package pool

import (
	"go-network-scanner/api"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	wg     = sync.WaitGroup{}
	called = atomic.Uint32{}
)

type workerMock struct {
	called bool
}

func (w *workerMock) Do() *api.ScannedResultsDTO {
	wg.Done()
	called.Add(1)
	w.called = true

	return &api.ScannedResultsDTO{}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name      string
		w         []Worker
		wantCalls int
	}{
		{
			"one worker",
			[]Worker{
				&workerMock{},
			},
			1,
		},
		{
			"three workers",
			[]Worker{
				&workerMock{},
				&workerMock{},
				&workerMock{},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(tt.wantCalls)
			called.Store(0)

			Start(NewHub(tt.w))

			wg.Wait()
			assert.Equal(t, tt.wantCalls, int(called.Load()))
		})
	}
}

func TestStartWithNoWorkers(t *testing.T) {
	noWorkers := []Worker{}
	hub := NewHub(noWorkers)
	assert.Panics(t, func() { hub.start() })
}
