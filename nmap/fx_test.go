package nmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScanner(t *testing.T) {
	ipRange := "ip1"
	proxy := "proxy1"
	scanner := NewScanner(ipRange, proxy)
	assert.NotNil(t, scanner)
}
