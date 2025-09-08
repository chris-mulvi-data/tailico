package test

import (
	"chris-mulvi-data/tailico/internal"
	"fmt"
	"testing"
)

func TestOutputLine(t *testing.T) {
	var tests = []string{
		"FATAL 18:20:33,272  [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an FATAL message from 127.0.0.1:50130",
		"18:20:34,530 ERROR [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an ERROR message from 127.0.0.1:50130",
		"18:20:35,332 WARN  [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an WARN message from 127.0.0.1:50130",
		"18:20:35,332 WARNING  [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an WARN message from 127.0.0.1:50130",
		"18:20:36,254 INFO  [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an INFO message from 127.0.0.1:50130",
		"18:20:37,156 DEBUG [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an DEBUG message from 127.0.0.1:50130",
		"18:20:38,003 TRACE [org.jboss.as.quickstarts.logging.rest.LogListingResource] (default task-1) This is an TRACE message from 127.0.0.1:50130",
	}

	for _, test := range tests {
		line := test
		internal.OutputLine(&line)
		fmt.Println(line)
	}
}
