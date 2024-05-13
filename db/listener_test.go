package db

import (
	"go-network-scanner/api"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type repoMock struct {
	callCount int
	wg        sync.WaitGroup

	upsertFn func(host *api.Host) (int, error)
}

func (m *repoMock) All() ([]*api.Host, error) {
	return nil, nil
}

func (m *repoMock) Upsert(h *api.Host) (int, error) {
	m.wg.Done()
	m.callCount++

	if m != nil && m.upsertFn != nil {
		return m.upsertFn(h)
	}
	return 1, nil
}

func (m *repoMock) currentCount() int {
	m.wg.Wait()
	return m.callCount
}

func TestHostTopic(t *testing.T) {
	l := NewDbHostListener(&repoMock{})
	assert.Equal(t, api.EventNmapHost, string(l.Topic()))
}

func TestHostListen(t *testing.T) {
	tests := []struct {
		name      string
		input     []any
		wantCalls int
	}{
		{
			name:      "success",
			input:     []any{&api.Host{}},
			wantCalls: 1,
		},
		{
			name:      "cast error",
			input:     []any{&api.Config{}},
			wantCalls: 0,
		},
		{
			name:      "empty input",
			input:     []any{},
			wantCalls: 0,
		},
		{
			name:      "nil input",
			input:     nil,
			wantCalls: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &repoMock{}
			mock.wg.Add(tt.wantCalls)

			l := NewDbHostListener(mock)
			l.Listen(tt.input...)

			assert.Equal(t, tt.wantCalls, mock.currentCount())
		})
	}
}

func TestHostProcess(t *testing.T) {
	tests := []struct {
		name    string
		mock    *repoMock
		wantErr bool
	}{
		{
			name:    "success",
			mock:    &repoMock{},
			wantErr: false,
		},
		{
			name: "upsert error",
			mock: &repoMock{
				upsertFn: func(*api.Host) (int, error) { return 0, assert.AnError },
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock.wg.Add(1)

			l := NewDbHostListener(tt.mock)
			id, err := l.process(&api.Host{})
			if tt.wantErr {
				assert.Equal(t, 0, id)
				assert.Error(t, err)
				assert.ErrorIs(t, err, assert.AnError)
			} else {
				assert.Equal(t, 1, id)
				assert.NoError(t, err)
			}
		})
	}
}
