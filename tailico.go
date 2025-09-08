// Tailico is a version of the tail command that includes
// colors when reading the output of server logs
package main

import (
	i "chris-mulvi-data/tailico/internal"
	"os"
)

func main() {
	opts, err := i.ParseArgs()
	if err != nil {
		i.OutputError(err.Error())
		// TODO: display "usage" tips
		return
	}

	if opts.Help {
		// TODO: display "usage" tips
		return
	}

	file, err := os.Open(opts.File)
	lines, _, err := i.ReadNLines(file, opts.LineCount)
	for _, line := range lines {
		i.OutputLine(&line)
	}
	file.Close()

	if opts.Tail {
		err := i.Tail(opts.File)
		if err != nil {
			i.OutputError(err.Error())
		}
	}

}
