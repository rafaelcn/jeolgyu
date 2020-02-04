package jeolgyu

const (
	// SinkFile outputs every sentence sent to the logger to a file
	SinkFile = 0x2
	// SinkOutput redirects every sentence sent to the logger to the stabdard
	// output
	SinkOutput = 0x4
	// SinkBoth redirects every
	SinkBoth = SinkFile | SinkOutput
)