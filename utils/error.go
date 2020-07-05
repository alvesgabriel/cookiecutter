package utils

import "log"

// FatalError print a log fatal and exit
func FatalError(err error) {
	if err != nil {
		log.Fatalf("ERROR: %#v", err.Error())
		log.Fatalf("TOTAL ERROR: %#v", err)
	}
}
