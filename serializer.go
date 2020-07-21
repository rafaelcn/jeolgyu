package jeolgyu

import (
	"encoding/json"
	"fmt"
)

const (
	stdoutInfo    string = " [+] "
	stdoutError   string = " [x] "
	stdoutWarning string = " [!] "
	stdoutPanic   string = " [X] "
)

// serializeToFile returns a marshalled message given the default output message
func serializeToFile(level Level, what, when string) []byte {
	output := make(MessageFormat, 1)

	output["level"] = level.string()
	output["what"] = what
	output["when"] = when

	message, err := json.Marshal(output)

	if err != nil {
		msg := fmt.Sprintf("Couldn't transform message to JSON. Reason %v", err)
		panic(msg)
	}

	return message
}

func serializeToOutput(level Level, what, when string) []byte {
	m := ""

	switch level {
	case InfoLevel:
		m = when + stdoutInfo + what
	case WarningLevel:
		m = when + stdoutWarning + what
	case ErrorLevel:
		m = when + stdoutError + what
	case PanicLevel:
		m = when + stdoutPanic + what
	}

	b := []byte(m)
	b = append(b, '\n')

	return b
}
