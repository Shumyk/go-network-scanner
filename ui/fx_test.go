package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUiListener(t *testing.T) {
	l := NewUIListener(nil)
	assert.NotNil(t, l)
}

func TestNewWebUI(t *testing.T) {
	ui := NewWebUI()
	assert.NotNil(t, ui)
}
