package jeolgyu

// Level abstracts the message level of a message
type Level int8

const (
	// InfoLevel is information message telling you about something.
	InfoLevel = iota + 1
	// WarningLevel message would be a thing to pay attetion whenever messages
	// of this type appears
	WarningLevel
	// ErrorLevel means that wrong happened, hence the error.
	ErrorLevel
	// FatalLevel is a so strong message you don't want it happening on your
	// system.
	FatalLevel
)

// Returns a string representation of the level
func (l Level) string() string {
	s := ""

	switch l {
	case InfoLevel:
		s = "info"
	case WarningLevel:
		s = "warning"
	case ErrorLevel:
		s = "error"
	case FatalLevel:
		s = "fatal"
	}

	return s
}
