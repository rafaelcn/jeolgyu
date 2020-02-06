package jeolgyu

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

// New creates a jeolgyu logger. The sink is given as one of the three:
// SinkFile, SinkOutput, SinkBoth. The fp parameter is only given if you want to
// specify where the loggerfile must be created, this parameter can be given as
// a relative path.
func New(sink int, fp string) (*Jeolgyu, error) {
	filename := ""
	var file *os.File
	var err error

	if (sink & SinkFile) == SinkFile {
		t := time.Now().Format("2006-Jan-2 15h 04m 05s")
		filename = t + ".log"

		if !exists(filename) {
			abs, _ := filepath.Abs(fp)
			f := path.Join(abs, filename)
			file, err = os.Create(f)

			if err != nil {
				const msg = "Error trying to create log file %s. Reason %v"
				e := fmt.Errorf(msg, filename, err)

				return nil, e
			}
		} else {
			file, err = os.Open(filename)

			if err != nil {
				const msg = "Error trying to open log file %s. Reason %v"
				e := fmt.Errorf(msg, filename, err)

				return nil, e
			}
		}
	}

	j := &Jeolgyu{
		sink:     sink,
		filename: filename,
		file:     file,
	}

	return j, nil
}

// Info prints information messages to whathever sink is selected
func (j *Jeolgyu) Info(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	serialized := Serialize(InfoLevel, m, t)

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

	serialized := Serialize(WarningLevel, m, t)

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

	serialized := Serialize(PanicLevel, m, t)

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

// This function exists with debugging/testing purposes to return the name of
// the file that Jeolgyu created when initialized.
func (j *Jeolgyu) getFilename() string {
	return j.filename
}

// This function exists with debugging/testing purposes to return the file that
// Jeolgyu created when initialized.
func (j *Jeolgyu) getFile() *os.File {
	return j.file
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
