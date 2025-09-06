package internal

import (
	"bufio"
	"container/ring"
	"os"
)

func ReadNLines(file *os.File, numLines int) ([]string, int64, error) {

	r := ring.New(numLines)
	scanner := bufio.NewScanner(file)
	var offset int64 = 0

	for scanner.Scan() {
		r.Value = scanner.Text()
		r = r.Next()
		offset += int64(len(scanner.Bytes())) + 1
	}

	var lines []string
	r.Do(func(a any) {
		if a != nil {
			lines = append(lines, a.(string))
		}
	})

	return lines, offset, scanner.Err()
}
