package pool

import (
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

func (w *workerMock) Do() {
	wg.Done()
	called.Add(1)
	w.called = true
}

func TestStart(t *testing.T) {
	tests := []struct {
		name      string
		w         []Worker
		wantCalls int
	}{
		{
			"no workers",
			[]Worker{},
			0,
		},
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
