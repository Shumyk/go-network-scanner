package nmap

import (
	"errors"
	"testing"

	"github.com/Ullaakut/nmap/v3"
	"github.com/stretchr/testify/assert"
)

type mockScanner struct {
	shouldError bool
}

func (m *mockScanner) Run() (result *nmap.Run, warnings *[]string, err error) {
	if m.shouldError {
		return nil, nil, errors.New("mocked error")
	}
	return &nmap.Run{}, &[]string{"warning1"}, nil
}

func TestScannerScan(t *testing.T) {
	tests := []struct {
		name    string
		scanner scanRunner
		panics  bool
	}{
		{
			name:    "nil scanner",
			scanner: nil,
			panics:  true,
		},
		{
			name: "scanner error free",
			scanner: &mockScanner{
				shouldError: false,
			},
			panics: false,
		},
		{
			name: "scanner with error",
			scanner: &mockScanner{
				shouldError: true,
			},
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scanner{scanner: tt.scanner}
			if tt.panics {
				assert.Panics(t, func() { s.Scan() })
			} else {
				assert.NotPanics(t, func() { s.Scan() })
				assert.NotNil(t, s.Scan())
			}
		})
	}
}

func TestScannerCheckError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		panics bool
	}{
		{
			name:   "nil error",
			err:    nil,
			panics: false,
		},
		{
			name:   "error",
			err:    errors.New("mocked error"),
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scanner{}
			if tt.panics {
				assert.Panics(t, func() { s.checkError(tt.err) })
			} else {
				assert.NotPanics(t, func() { s.checkError(tt.err) })
			}
		})
	}
}
