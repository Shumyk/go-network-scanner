package db

import (
	log "log/slog"

	"github.com/kataras/go-events"
	"github.com/pkg/errors"
)

const zero = 0

// Listener listens for events and persists them.
type Listener[T any] struct {
	repo  RepoBasic[T]
	topic events.EventName
}

// Topic returns the event topic.
func (l *Listener[T]) Topic() events.EventName {
	return l.topic
}

// Listen listens for events.
func (l *Listener[T]) Listen(i ...any) {
	if len(i) == zero || i[zero] == nil {
		log.Error("no event data")
		return
	}
	h, ok := i[zero].(T)
	if !ok {
		log.Error("failed to cast")
		return
	}
	//nolint:errcheck // Ignore error here as it's returned for tests.
	go l.process(h)
}

func (l *Listener[T]) process(t T) (int, error) {
	id, err := l.repo.Upsert(t)
	if err != nil {
		log.Error("failed to persist in DB", log.Any("error", err))
		return id, errors.Wrap(err, "failed to persist in DB")
	}

	log.Debug(
		"persisted",
		log.Any("id", id),
		log.Any("data", t),
		log.Any("topic", l.topic),
	)
	return id, nil
}
