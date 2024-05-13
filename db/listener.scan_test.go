package db

import (
	"go-network-scanner/api"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type scanRepoMock struct {
	callCount int
	wg        sync.WaitGroup

	upsertFn func(h *api.Scan) (int, error)
}

func (m *scanRepoMock) All() ([]*api.Scan, error) {
	return nil, nil
}

func (m *scanRepoMock) Upsert(h *api.Scan) (int, error) {
	m.wg.Done()
	m.callCount++

	if m != nil && m.upsertFn != nil {
		return m.upsertFn(h)
	}
	return 1, nil
}

func (m *scanRepoMock) currentCount() int {
	m.wg.Wait()
	return m.callCount
}

func TestScanTopic(t *testing.T) {
	var l *ScanListener
	assert.Equal(t, api.EventNmapScan, string(l.Topic()))
}

func TestScanListen(t *testing.T) {
	tests := []struct {
		name      string
		input     []any
		mock      *scanRepoMock
		wantCalls int
	}{
		{
			name:      "success",
			input:     []any{&api.Scan{}},
			mock:      &scanRepoMock{},
			wantCalls: 1,
		},
		{
			name:      "cast error",
			input:     []any{&api.Config{}},
			mock:      &scanRepoMock{},
			wantCalls: 0,
		},
		{
			name:      "no input",
			input:     []any{},
			mock:      &scanRepoMock{},
			wantCalls: 0,
		},
		{
			name:      "nil input",
			input:     []any{nil},
			mock:      &scanRepoMock{},
			wantCalls: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock.wg.Add(tt.wantCalls)

			l := NewDbScanListener(tt.mock)
			l.Listen(tt.input...)

			assert.Equal(t, tt.wantCalls, tt.mock.currentCount())
		})
	}
}

func TestScanProcess(t *testing.T) {
	tests := []struct {
		name      string
		mock      *scanRepoMock
		wantError bool
	}{
		{
			name: "success",
			mock: &scanRepoMock{},
		},
		{
			name: "error",
			mock: &scanRepoMock{
				upsertFn: func(h *api.Scan) (int, error) { return 0, assert.AnError },
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock.wg.Add(1)

			l := &ScanListener{tt.mock}
			id, err := l.process(&api.Scan{})
			if tt.wantError {
				assert.Equal(t, 0, id)
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, assert.AnError)
			} else {
				assert.Equal(t, 1, id)
				assert.NoError(t, err)
			}
		})
	}
}
