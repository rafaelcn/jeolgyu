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

		if !exists(filename) {
			abs, _ := filepath.Abs(s.Filepath)
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
		sink:     s.SinkType,
		filename: filename,
		file:     file,
	}

	return j, nil
}

// Info prints information messages to whathever sink is selected
func (j *Jeolgyu) Info(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	j.sinkTo(Serialize(InfoLevel, m, t))
}

// Warning prints a warning message to whatever sink is selected
func (j *Jeolgyu) Warning(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	j.sinkTo(Serialize(WarningLevel, m, t))
}

// Error prints an error message to whatever sink is selected
func (j *Jeolgyu) Error(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	j.sinkTo(Serialize(ErrorLevel, m, t))
}

// Panic prints a message to whatever sink is selected
func (j *Jeolgyu) Panic(message string, arguments ...interface{}) {
	t := now()
	m := format(message, arguments...)

	j.sinkTo(Serialize(PanicLevel, m, t))
}

// sinkTo sends the message to whatever sink j is set to
func (j *Jeolgyu) sinkTo(m []byte) {
	switch j.sink {
	case SinkBoth:
		sinkOutput(m)
		j.sinkFile(m, j.file)
	case SinkFile:
		j.sinkFile(m, j.file)
	case SinkOutput:
		sinkOutput(m)
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

func (j *Jeolgyu) sinkFile(message []byte, file *os.File) {
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
