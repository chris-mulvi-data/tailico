package internal

import (
	"fmt"
)

func OutputError(msg string) {
	fmt.Printf("%s%s%s\n", Red, msg, Default)
}

func OutputMsg(msg string) {
	fmt.Printf("%s%s%s\n", Cyan, msg, Default)
}

func ColorizeByType(line *string) {
	// TODO: check if a timestamp is at the front of the line
	// TODO: extract the timestamp if it is there
	// TODO: colorize the rest of the line
	// TODO: stick the timestamp back on
}
