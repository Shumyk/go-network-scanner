package ui

import _ "embed"

// Embedded index.html file. This file contains build UI into one file.
//
//go:embed pages/build/index.html
var page string
