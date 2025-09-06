package internal

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Tail(file *os.File, offset int64) error {
	fmt.Println("... starting tail...")
	fmt.Printf("offset: %d", offset)

	file.Seek(offset, 0)
	reader := bufio.NewReader(file)

	// monitor the file for new lines
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// TODO: call the formatting functions here
		fmt.Print(line)
		continue
	}
}
