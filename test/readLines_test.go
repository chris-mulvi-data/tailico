package test

import (
	"chris-mulvi-data/tailico/internal"
	"fmt"
	"os"
	"testing"
)

func TestReadNLines(t *testing.T) {
	file, err := os.Open("../testing.txt")
	if err != nil {
		t.Error(err)
	}
	lines, offset, err := internal.ReadNLines(file, 5)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("offset: %d\n", offset)
	for _, line := range lines {
		fmt.Println(line)
	}
}
