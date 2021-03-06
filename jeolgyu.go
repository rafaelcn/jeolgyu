package jeolgyu

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

type (
	// Jeolgyu encapsulates the settings of the logger
	Jeolgyu struct {
		sink     Sink
		filename string
		file     *os.File
		// write file lock
		mu sync.Mutex
	}

	// Settings encapsulates the settings for the logger and it is
	// accepted in its constructor.
	Settings struct {
		// SinkType specifies the type of the sink used by the initialized
		// logger. There are currently three types of sinks: SinkFile,
		// SinkOutput and SinkBoth.
		//
		// Whenever you use the SinkFile a file containing the log content will
		// be created with a timestamp of creation of the specified filename.
		// Using the SinkOutput the logger will direct its messages to the
		// stdout or stderr, depending on the level of the message.
		SinkType Sink
		// Filepath is an optional field that indicates where the logger file
		// should be created. Giving it an empty string will lead Jeolgyu to
		// create in the current running directory.
		Filepath string
		// Filename is an optional field that sets the name of the log file.
		Filename string
	}

	// MessageFormat is a type containing the possible format for messages and
	// is defined in the serializer.
	MessageFormat map[string]string
)

// New creates a jeolgyu logger with a settings struct
func New(s Settings) (*Jeolgyu, error) {
	filename := ""
	var file *os.File
	var err error

	if (s.SinkType & SinkFile) == SinkFile {
		t := ""

		// create a file with a timestamp on its name only if the user did not
		// specified the name on creation.
		if len(s.Filename) == 0 {
			t = time.Now().Format("2006-Jan-2 15h 04m 05s")
			filename = t + ".log"
		} else {
			filename = s.Filename + ".log"
		}

		// create the specified folder if it doesn't exist
		if !exists(s.Filepath) && len(s.Filepath) > 0 {
			err := os.Mkdir(s.Filepath, 0744)

			if err != nil {
				e := fmt.Errorf("Error trying to create folder %s. Reason %v",
					s.Filepath, err)

				return nil, e
			}
		}

		abs, _ := filepath.Abs(s.Filepath)
		filename = path.Join(abs, filename)

		if !exists(filename) {
			file, err = os.Create(filename)

			if err != nil {
				const msg = "Error trying to create log file %s. Reason %v"
				e := fmt.Errorf(msg, filename, err)

				return nil, e
			}
		} else {
			file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

			if err != nil {
				const msg = "Error trying to open log file %s. Reason %v"
				e := fmt.Errorf(msg, filename, err)

				return nil, e
			}
		}
	}

	j := &Jeolgyu{
		sink:     s.SinkType,
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

// Fatal prints a message to whatever sink is selected and exits
func (j *Jeolgyu) Fatal(message string, arguments ...interface{}) {
	j.sinkTo(FatalLevel, message, arguments...)
	os.Exit(1)
}

// sinkTo sends the message to whatever sink j is set to
func (j *Jeolgyu) sinkTo(level Level, message string, arguments ...interface{}) {
	m := format(message, arguments...)

	switch j.sink {
	case SinkBoth:
		j.sinkOutput(level, m)
		j.sinkFile(level, m, j.file)
	case SinkFile:
		j.sinkFile(level, m, j.file)
	case SinkOutput:
		j.sinkOutput(level, m)
	}
}

// sinkOutput formats the message to the stdout
func (j *Jeolgyu) sinkOutput(level Level, message string) {
	t := time.Now().Format("15:04:05")

	m := serializeToOutput(level, message, t)

	if level == ErrorLevel {
		fmt.Fprint(os.Stderr, string(m))
	} else {
		fmt.Fprint(os.Stdout, string(m))
	}
}

// sinkFile appends a message to the current file log with a serialized output
func (j *Jeolgyu) sinkFile(level Level, message string, file *os.File) {
	j.mu.Lock()
	defer j.mu.Unlock()

	t := time.Now().Format("2006-Jan-2 15:04:05")

	m := serializeToFile(level, message, t)
	m = append(m, '\n')

	file.Write(m)
}
