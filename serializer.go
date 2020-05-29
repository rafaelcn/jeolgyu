package jeolgyu

import (
	"encoding/json"
	"fmt"
)

// serialize returns a marshalled message given the default output message
func serialize(level Level, what, when string) []byte {
	output := make(MessageFormat, 1)

	output["level"] = level.String()
	output["what"] = what
	output["when"] = when

	message, err := json.Marshal(output)

	if err != nil {
		msg := fmt.Sprintf("Couldn't transform message to JSON. Reason %v", err)
		panic(msg)
	}

	return message
}
