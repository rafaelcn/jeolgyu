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
	// PanicLevel is a so strong message you don't want it happening on your
	// system.
	PanicLevel
)

// Returns a string representation of the level
func (l Level) String() string {
	s := ""

	switch l {
	case InfoLevel:
		s = "info"
	case WarningLevel:
		s = "warning"
	case ErrorLevel:
		s = "error"
	case PanicLevel:
		s = "panic"
	}

	return s
}
