package internal

import "regexp"

// FindTimestampInString finds a timestamp in the pointer of a string.
//
// Using a pointer as the input to limit duplicating memory use
func FindTimestampInString(str *string) string {
	formats := getFormats()
	var timestamp string
	for _, v := range formats {
		re := regexp.MustCompile(v)
		match := re.FindString(*str)
		if match == "" {
			continue
		}
		timestamp = match
		break
	}
	return timestamp
}

// getFormats gets the timestamp formats map
func getFormats() map[string]string {

	// formats must me listed in order of most specific to least specific
	var formats = make(map[string]string)
	formats["yyyy-mm-dd hh:mm:ss.SSS"] = "\\d{4}-\\d{2}-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}\\.\\d{3}"
	formats["yyyy-mm-dd hh:mm:ss,SSS"] = "\\d{4}-\\d{2}-\\d{2}\\s\\d{2}:\\d{2}:\\d{2},\\d{3}"
	formats["yyyy-mm-dd hh:mm:ss"] = "\\d{4}-\\d{2}-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
	formats["HH:mm:ss,SSS"] = "\\d{2}:\\d{2}:\\d{2},\\d{3}"

	return formats
}
