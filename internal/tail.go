package internal

import (
	"bufio"
	"os/exec"
)

func Tail(filename string) error {
	cmd := exec.Command("tail", "-f", filename)
	// Create a pipe to read the output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Create a buffered reader
	reader := bufio.NewReader(stdout)

	// Continuously read lines
	for {
		// Read until newline
		line, err := reader.ReadString('\n')
		if err != nil {
			OutputError(err.Error())
			break
		}

		// Process the line
		OutputLine(&line)
	}

	// Wait for the command to complete
	return cmd.Wait()

}
