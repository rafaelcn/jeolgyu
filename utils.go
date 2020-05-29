package jeolgyu

import (
	"fmt"
	"os"
	"time"
)

func exists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}

	return true
}

func now() string {
	return time.Now().Format("2006-Jan-2 15h 04m 05s")
}

func format(message string, arguments ...interface{}) string {
	if arguments == nil || len(arguments) == 0 {
		return message
	}

	return fmt.Sprintf(message, arguments...)
}
