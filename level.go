package jeolgyu

type Level int8

const (
	// An information message telling you about something. 
	InfoLevel = iota + 1
	// A warning level message. If I'd describe it, it would be a thing to pay
	// attetion whenever messages of this type appears
	WarningLevel
	// Something wrong happened, hence the error.
	ErrorLevel
	// A mesage level so strong you don't want it happening on yout system.
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
	case PanicLevel:
		s = "panic"
	}

	return s
}
