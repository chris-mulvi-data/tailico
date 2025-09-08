package test

import (
	i "chris-mulvi-data/tailico/internal"
	"os"
	"testing"
)

// TestTail is a test to see if the function will read new lines
// works best in a debug
func TestTail(t *testing.T) {
	file, err := os.Open("../testing.txt")
	if err != nil {
		t.Error(err)
	}

	i.Tail(file)
}
