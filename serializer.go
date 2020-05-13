package jeolgyu

import (
	"encoding/json"
	"fmt"
)

// serializeToFile returns a marshalled message given the default output message
func serializeToFile(level Level, what, when string) []byte {
	output := make(MessageFormat, 1)

	output["level"] = level.String()
	output["what"] = what
	output["when"] = when

	message, err := json.Marshal(output)

	if err != nil {
		msg := fmt.Sprintf("Couldn't transform message to JSON. Reason %v", err)
		panic(msg)
	}

	return message
}

func serializeToOutput(level Level, what, when string) string {
	message := ""

	switch level {
	case InfoLevel:
		message = when + " [+] "
		message = message + what
	case WarningLevel:
		message = when + " [!] "
		message = message + what
	case ErrorLevel:
		message = when + " [x] "
		message = message + what
	case PanicLevel:
		message = when + " [P] "
		message = message + what
	}

	return message
}