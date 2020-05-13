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
		sink     Sink
		filename string
		file     *os.File
	}

	// MessageFormat is a type containing the possible format for messages and
	// is defined in the serializer.
	MessageFormat map[string]string
)

// New creates a jeolgyu logger. The sink is given as one of the three:
// SinkFile, SinkOutput, SinkBoth.
//
// The fp (filepath) parameter is only given if you want to specify where the
// loggerfile must be created, this parameter can be given as a relative path,
// by default it creates log files under the same directory of the running
// application.
func New(sink Sink, fp string) (*Jeolgyu, error) {
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
	j.sinkTo(InfoLevel, message, arguments...)
}

// Warning prints a warning message to whatever sink is selected
func (j *Jeolgyu) Warning(message string, arguments ...interface{}) {
	j.sinkTo(WarningLevel, message, arguments...)
}

// Error prints an error message to whatever sink is selected
func (j *Jeolgyu) Error(message string, arguments ...interface{}) {
	j.sinkTo(ErrorLevel, message, arguments...)
}

// Panic prints a message to whatever sink is selected
func (j *Jeolgyu) Panic(message string, arguments ...interface{}) {
	j.sinkTo(PanicLevel, message, arguments...)
}

// sinkTo sends the message to whatever sink j is set to
func (j *Jeolgyu) sinkTo(level Level, message string, arguments ...interface{}) {
	m := format(message, arguments...)

	switch j.sink {
	case SinkBoth:
		sinkOutput(level, m)
		sinkFile(level, m, j.file)
	case SinkFile:
		sinkFile(level, m, j.file)
	case SinkOutput:
		sinkOutput(level, m)
	}
}

// sinkOutput formats the message to the stdout
func sinkOutput(level Level, message string) {
	t := time.Now().Format("15:04:05")

	fmt.Println(serializeToOutput(level, message, t))
}

// sinkFile appends a message to the current file log with a serialized output
func sinkFile(level Level, message string, file *os.File) {
	t := time.Now().Format("2006-Jan-2 15:04:05")
	m := serializeToFile(level, message, t)

	m = append(m, '\n')
	file.WriteString(string(m))
}

// getFilename exists with debugging/testing purposes to return the name of
// the file that Jeolgyu created when initialized.
func (j *Jeolgyu) getFilename() string {
	return j.filename
}

// getFile exists with debugging/testing purposes to return the file that
// Jeolgyu created when initialized.
func (j *Jeolgyu) getFile() *os.File {
	return j.file
}
