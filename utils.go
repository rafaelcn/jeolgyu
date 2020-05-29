package jeolgyu

import (
	"fmt"
	"os"
<<<<<<< HEAD
)

// exists verify the existence of a file and returns true if the file exists and
// false otherwise
=======
	"time"
)

>>>>>>> Make if thread safe and remove unnecessary funs
func exists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}

	return true
}

<<<<<<< HEAD
=======
func now() string {
	return time.Now().Format("2006-Jan-2 15h 04m 05s")
}

>>>>>>> Make if thread safe and remove unnecessary funs
func format(message string, arguments ...interface{}) string {
	if arguments == nil || len(arguments) == 0 {
		return message
	}

	return fmt.Sprintf(message, arguments...)
}
