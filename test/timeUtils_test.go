package test

import (
	"chris-mulvi-data/tailico/internal"
	"fmt"
	"testing"
)

func TestFindTimestampInString(t *testing.T) {
	var testStrings = []string{
		"2025-09-01 12:10:30",
		"2025-09-01 12:10:30.000",
		"2025-09-01 12:10:30,000",
		"18:20:34,530",
	}

	for _, str := range testStrings {
		timestamp := internal.FindTimestampInString(&str)
		if timestamp != str {
			t.Errorf("No match: %s -> %s\n", str, timestamp)
			continue
		}
		fmt.Printf("%s -> %s\n", str, timestamp)

	}

}
