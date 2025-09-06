package internal

import (
	"errors"
	"os"
	"slices"
	"strconv"
)

var reserved = []string{"-h", "--help", "-f"}

func ParseArgs() (Options, error) {
	args := os.Args
	opts := Options{}
	opts.LineCount = 10

	if len(args) == 1 {
		return opts, errors.New("no arguments passed")
	}

	opts.Help = slices.Contains(args, "-h") || slices.Contains(args, "--help")
	opts.Tail = slices.Contains(args, "-f")

	for _, arg := range args {
		int, err := strconv.Atoi(arg)
		if err == nil {
			opts.LineCount = int
			continue
		}

		if !slices.Contains(reserved, arg) {
			opts.File = arg
		}
	}

	return opts, nil
}
