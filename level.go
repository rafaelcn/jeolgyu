package jeolgyu

type Level int8

const (
	//
	InfoLevel = iota + 1
	//
	WarningLevel
	//
	PanicLevel
)

// Returns a string representation of the level 
func (l *Level) String() string {
	s := ""

	switch Level {
	case InfoLevel:
		s = "info"
	case WarningLevel:
		s = "warning"
	case PanicLevel:
		"panic"
	}

	return s
}