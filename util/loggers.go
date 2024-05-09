package util

import (
	log "log/slog"
	"os"
	"path/filepath"
)

// SetupLogger sets up the default logger for the application.
func SetupLogger() {
	log.SetDefault(Logger())
}

// Logger returns a new logger with the default handler.
func Logger() *log.Logger {
	return log.New(handler())
}

// LoggerWith returns a new logger with the default handler
// and the given group and attributes.
func LoggerWith(group string, attrs ...log.Attr) *log.Logger {
	return log.New(
		handler().
			WithGroup(group).
			WithAttrs(attrs),
	)
}

func handler() log.Handler {
	return log.NewTextHandler(
		os.Stdout,
		&log.HandlerOptions{
			AddSource:   true,
			Level:       log.LevelInfo,
			ReplaceAttr: replaceFullSourcePath,
		},
	)
}

func replaceFullSourcePath(_ []string, attr log.Attr) log.Attr {
	if attr.Key == log.SourceKey {
		source, ok := attr.Value.Any().(*log.Source)
		if ok {
			// directory/filename.go
			source.File = filepath.Join(
				filepath.Base(filepath.Dir(source.File)),
				filepath.Base(source.File),
			)
		}
	}
	return attr
}
