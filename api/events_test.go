package api

import (
	"testing"

	"github.com/kataras/go-events"
	"github.com/stretchr/testify/assert"
)

type listener struct {
	topic events.EventName
}

func (l listener) Topic() events.EventName {
	return l.topic
}

func (l listener) Listen(i ...any) {}

func TestSubscribeListeners(t *testing.T) {
	tests := []struct {
		name string
		ll   []Listener
		want int
	}{
		{
			name: "empty",
			ll:   []Listener{},
			want: 0,
		},
		{
			name: "one listener",
			ll: []Listener{
				&listener{"cool.topic"},
			},
			want: 1,
		},
		{
			name: "two listeners",
			ll: []Listener{
				&listener{"some.topic"},
				&listener{"another.topic"},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events.Clear()
			SubscribeListeners(tt.ll)
			assert.Equal(t, countListeners(), tt.want)
		})
	}
}

func countListeners() (c int) {
	for _, n := range events.EventNames() {
		c += events.ListenerCount(n)
	}
	return
}
