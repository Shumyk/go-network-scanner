package util

import (
	"go-network-scanner/api"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskInput(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		want   api.Config
		wantDb string
	}{
		{
			name:  "proper input",
			input: []string{"session1", "local_db", "192.33.55.31", "192.32.55.59", "192.123.43.64", "y"},
			want: api.Config{
				Name:        "session1",
				IpRange1:    "192.33.55.31",
				IpRange2:    "192.32.55.59",
				IpRange3:    "192.123.43.64",
				EnableProxy: true,
			},
			wantDb: "local_db",
		},
		{
			name:  "skipping all optional fields",
			input: []string{"session2", "", "", "33.33.55.31", "", "", ""},
			want: api.Config{
				Name:        "session2",
				IpRange1:    "33.33.55.31",
				IpRange2:    "",
				IpRange3:    "",
				EnableProxy: false,
			},
			wantDb: DEFAULT_TEST_DB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input = createInput(t, tt.input)
			defer input.Close()

			databaseName, config := AskInput()
			assert.Equal(t, tt.want, config)
			assert.Equal(t, tt.wantDb, databaseName)
		})
	}
}

func TestAskDatabaseName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "use default", input: "", want: DEFAULT_TEST_DB},
		{name: "proper input", input: "this", want: "this"},
		{name: "extra spaces - trim", input: " that ", want: "that"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input = createInput(t, []string{tt.input})
			defer input.Close()

			assert.Equal(t, tt.want, askDatabaseName())
		})
	}
}

func TestAskSessionName(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "on first try", input: []string{"this", "that"}, want: "this"},
		{name: "on second try", input: []string{"", "second"}, want: "second"},
		{name: "on third try", input: []string{"", "", "pip"}, want: "pip"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input = createInput(t, tt.input)
			defer input.Close()

			config := api.Config{}
			askSessionName(&config)
			assert.Equal(t, tt.want, config.Name)
		})
	}
}

func TestAskIpRange(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "on first try", input: []string{"44.32.55.31", "127.0.0.1"}, want: "44.32.55.31"},
		{name: "on second try", input: []string{"", "34.32.55.31"}, want: "34.32.55.31"},
		{name: "on third try", input: []string{"", "", "192.32.55.31"}, want: "192.32.55.31"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input = createInput(t, tt.input)
			defer input.Close()

			config := api.Config{}
			askIpRange(&config)
			assert.Equal(t, tt.want, config.IpRange1)
		})
	}
}

func TestAskSocksProxyEnabled(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  bool
	}{
		{
			name:  "positive",
			input: []string{"y", "yes", "YEAH!", " YEAH!", "  y  "},
			want:  true,
		},
		{
			name:  "negative",
			input: []string{"n", "no", "nO", "n0", "000n", "never", "NO!", " NO!", "  n  ", "", "!23", " what? ", " I love potatoes"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range tt.input {
				input = createInput(t, []string{i})
				defer input.Close()

				config := api.Config{}
				askSocksProxyEnabled(&config)
				assert.Equal(t, tt.want, config.EnableProxy)
			}
		})
	}
}

func createInput(t *testing.T, input []string) *os.File {
	in, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	for _, i := range input {
		_, err = io.WriteString(in, i+"\n")
		if err != nil {
			t.Fatal(err)
		}
	}
	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}
	return in
}
