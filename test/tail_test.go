package test

import (
	i "chris-mulvi-data/tailico/internal"
	"testing"
)

// TestTail is a test to see if the function will read new lines
// works best in a debug
func TestTail(t *testing.T) {
	filename := "../testing.txt"

	i.Tail(filename)
}
