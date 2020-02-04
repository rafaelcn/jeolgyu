package jeolgyu

import (
	"fmt"
	"os"
	"time"
)

type (
	// Jeolgyu encapsulates the settings of the logger
	Jeolgyu struct {
		sink     int
		filename string
		file     *os.File
	}

	// MessageFormat is a type containing the possible format for messages and
	// is defined in the serializer.
	MessageFormat map[string]string
)

// New creates a jeolgyu logger
func New(sink int) *Jeolgyu {
	filename := ""
	var file *os.File
	var err error

	if (sink & SinkFile) == 0x2 {
		t := time.Now().Format("2006-Jan-2 15h 04m 05s")
		filename = t + " .log"

		if !exists(filename) {
			file, err = os.Create(filename)

			if err != nil {
				const msg = "Error trying to create log file %s. Reason %v"
				panic(fmt.Sprintf(msg, filename, err))
			}
		} else {
			file, err = os.Open(filename)

			if err != nil {
				const msg = "Error trying to open log file %s. Reason %v"
				panic(fmt.Sprintf(msg, filename, err))
			}
		}
	}

	j := &Jeolgyu{
		sink:     sink,
		filename: filename,
		file:     file,
	}

	return j
}

// Info prints information messages to whathever sink is selected
func (j *Jeolgyu) Info(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	serialized := Serialize("info", m, t)

	switch j.sink {
	case SinkBoth:
		sinkOutput(serialized)
		sinkFile(serialized, j.file)
	case SinkFile:
		sinkFile(serialized, j.file)
	case SinkOutput:
		sinkOutput(serialized)
	}
}

// Warning prints a warning message to whatever sink is selected
func (j *Jeolgyu) Warning(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	serialized := Serialize("info", m, t)

	switch j.sink {
	case SinkBoth:
		sinkOutput(serialized)
		sinkFile(serialized, j.file)
	case SinkFile:
		sinkFile(serialized, j.file)
	case SinkOutput:
		sinkOutput(serialized)
	}
}

// Panic prints a message to whatever sink is selected and panics afterwards
func (j *Jeolgyu) Panic(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	serialized := Serialize("info", m, t)

	switch j.sink {
	case SinkBoth:
		sinkOutput(serialized)
		sinkFile(serialized, j.file)
	case SinkFile:
		sinkFile(serialized, j.file)
	case SinkOutput:
		sinkOutput(serialized)
	}
}

func sinkOutput(message []byte) {
	fmt.Print(string(message))
}

func sinkFile(message []byte, file *os.File) {
	message = append(message, '\n')
	file.WriteString(string(message))
}

func exists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return true
	}

	return false
}

func now() string {
	return time.Now().Format("2006-Jan-2 15h 04m 05s")
}

func format(message string, arguments ...interface{}) string {
	return fmt.Sprintf(message, arguments...)
}