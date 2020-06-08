package jeolgyu

import (
	"fmt"
	"os"
)

// exists verify the existence of a file and returns true if the file exists and
// false otherwise
func exists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}

	return true
}

func format(message string, arguments ...interface{}) string {
	if arguments == nil || len(arguments) == 0 {
		return message
	}

	return fmt.Sprintf(message, arguments...)
}
