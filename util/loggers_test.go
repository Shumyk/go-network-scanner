package util

import (
	"bufio"
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupLogger(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	writer = w
	SetupLogger()

	slog.Error("errors-many")
	w.Flush()

	parts := parseLog(b)
	checkMsg(t, parts, "errors-many")
	checkLevel(t, parts, slog.LevelError)
	checkSource(t, parts, "util/loggers_test.go:19")
}

func TestSourcePath(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	writer = w
	l := Logger()

	l.Info("Memes")
	w.Flush()

	parts := parseLog(b)
	checkMsg(t, parts, "Memes")
	checkLevel(t, parts, slog.LevelInfo)
	checkSource(t, parts, "util/loggers_test.go:34")
}

func TestLoggerWith(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	writer = w

	l := LoggerWith("worker", slog.Int("id", 15))
	l.Warn("working")
	w.Flush()

	parts := parseLog(b)
	checkMsg(t, parts, "working")
	checkLevel(t, parts, slog.LevelWarn)
	checkSource(t, parts, "util/loggers_test.go:49")

	// Check the group
	g, ok := parts["worker"]
	assert.True(t, ok)
	gm := g.(map[string]any)
	idI, ok := gm["id"]
	assert.True(t, ok)
	id, ok := idI.(string)
	assert.True(t, ok)
	assert.Equal(t, "15", id)
}

func parseLog(b bytes.Buffer) map[string]any {
	var m map[string]any = make(map[string]any)

	for _, line := range bytes.Split(b.Bytes(), []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(string(line), " ")
		for _, part := range parts {
			kv := strings.Split(part, "=")
			key := kv[0]
			val := kv[1]

			if strings.Contains(key, ".") {
				groupAttr := strings.Split(key, ".")
				group := groupAttr[0]
				attr := groupAttr[1]
				g, ok := m[group]
				if ok {
					gm := g.(map[string]any)
					gm[attr] = val
				} else {
					gm := make(map[string]any)
					gm[attr] = val
					m[group] = gm
				}
			} else {
				m[key] = val
			}

		}
	}
	return m
}

func checkSource(t *testing.T, parts map[string]any, expected string) {
	sourceI, ok := parts["source"]
	assert.True(t, ok)
	source, ok := sourceI.(string)
	assert.True(t, ok)
	assert.Equal(t, expected, source)
}

func checkLevel(t *testing.T, parts map[string]any, expected slog.Level) {
	levelI, ok := parts[slog.LevelKey]
	assert.True(t, ok)
	level, ok := levelI.(string)
	assert.True(t, ok)
	assert.Equal(t, expected.String(), level)
}

func checkMsg(t *testing.T, parts map[string]any, expected string) {
	msgI, ok := parts[slog.MessageKey]
	assert.True(t, ok)
	msg, ok := msgI.(string)
	assert.True(t, ok)
	assert.Equal(t, expected, msg)
}

func checkGroup(t *testing.T, parts map[string]any, group, attr, expected string) {
	g, ok := parts[group]
	assert.True(t, ok)
	gm := g.(map[string]any)
	idI, ok := gm[attr]
	assert.True(t, ok)
	id, ok := idI.(string)
	assert.True(t, ok)
	assert.Equal(t, expected, id)
}
