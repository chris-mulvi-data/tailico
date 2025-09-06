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

	var foremats = make(map[string]string)
	foremats["yyyy-mm-dd hh:mm:ss"] = ""

	return foremats
}
