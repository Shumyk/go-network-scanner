package ui

import (
	"embed"
	"io/fs"
	log "log/slog"

	"github.com/pkg/errors"
)

const indexHTML = "index.html"

// Embedded index.html file. This file contains build UI into one file.
//
//go:embed pages/build/index.html
var page embed.FS

func indexPage() string {
	indexPage, err := fs.ReadFile(page, indexHTML)
	if err != nil {
		log.Error("Failed to read index.html", log.Any("err", err))
		panic(errors.Wrap(err, "failed to read index.html"))
	}
	return string(indexPage)
}
