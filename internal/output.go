package internal

import (
	"fmt"
	"regexp"
	"strings"
)

var logTypePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(FATAL)\s+`),
	regexp.MustCompile(`(ERROR)\s+`),
	regexp.MustCompile(`(WARNING)\s+`),
	regexp.MustCompile(`(WARN)\s+`),
	regexp.MustCompile(`(DEBUG)\s+`),
	regexp.MustCompile(`(INFO)\s+`),
	regexp.MustCompile(`(TRACE)\s+`),
}

func OutputError(msg string) {
	fmt.Printf("%s%s%s\n", Red, msg, Default)
}

func OutputMsg(msg string) {
	fmt.Printf("%s%s%s\n", Cyan, msg, Default)
}

func OutputLine(line *string) {

	timestamp := FindTimestampInString(line)

	if timestamp != "" {
		// compile regex pattern using the identified timestamp
		re := regexp.MustCompile(timestamp)
		// colorize the timestamp
		timestamp = fmt.Sprintf("%s%s%s", Cyan, timestamp, Default)
		// replace the timestamp in the line with the colorized version
		*line = re.ReplaceAllString(*line, timestamp)
	}

	ColorizeByType(line)
	fmt.Print(*line)
}

func ColorizeByType(line *string) {
	for _, pattern := range logTypePatterns {
		matches := pattern.FindStringSubmatch(*line)
		if len(matches) < 1 {
			continue
		}
		var match string
		switch len(matches) {
		case 1:
			match = matches[0]
		case 2:
			match = matches[1]
		}

		color := getColor(match)
		match = fmt.Sprintf("%s%s%s ", color, match, Default)
		*line = pattern.ReplaceAllString(*line, match)
		return
	}
}

func getColor(str string) Color {

	switch strings.ToLower(str) {
	case "fatal":
		fallthrough
	case "error":
		return Red
	case "warn":
		fallthrough
	case "warning":
		return Yellow
	case "info":
		return Cyan
	case "debug":
		return Blue
	case "trace":
		return Gray
	}

	return Default
}
