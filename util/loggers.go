package util

import (
	"log/slog"
	"os"
	"path/filepath"
)

func SetupLogger() {
	slog.SetDefault(Logger())
}

func Logger() *slog.Logger {
	return slog.New(handler())

}

func LoggerWith(group string, attrs ...slog.Attr) *slog.Logger {
	return slog.New(
		handler().
			WithGroup(group).
			WithAttrs(attrs),
	)
}

func handler() slog.Handler {
	return slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replaceFullSourcePath,
		},
	)
}

func replaceFullSourcePath(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.SourceKey {
		source, _ := a.Value.Any().(*slog.Source)
		if source != nil {
			// directory/filename.go
			source.File = filepath.Join(
				filepath.Base(filepath.Dir(source.File)),
				filepath.Base(source.File),
			)
		}
	}
	return a
}
