package jeolgyu

// Sink is a type abstracting over the possible sinks of the logger
type Sink int8

const (
	// SinkFile outputs every sentence sent to the logger to a file
	SinkFile Sink = 0x2
	// SinkOutput redirects every sentence sent to the logger to the stabdard
	// output
	SinkOutput Sink = 0x4
	// SinkBoth redirects every
	SinkBoth Sink = SinkFile | SinkOutput
)
