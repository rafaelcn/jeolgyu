package jeolgyu

import (
	"reflect"
	"testing"
)

type serializerTests struct {
	Level      Level
	Message    string
	Arguments  []interface{}
	Want       string
	ShouldFail bool
}

func TestSerialize(t *testing.T) {
	tests := []serializerTests{
		{
			Level:      ErrorLevel,
			Message:    "A message with %s.",
			Arguments:  []interface{}{"arguments"},
			Want:       `{"level":"error","what":"A message with arguments.","when":""}`,
			ShouldFail: false,
		},
		{
			Level:      WarningLevel,
			Message:    "A message with no arguments.",
			Arguments:  nil,
			Want:       `{"level":"warning","what":"A message with no arguments.","when":""}`,
			ShouldFail: false,
		},
		{
			Level:      PanicLevel,
			Message:    "A %s message with %d %s.",
			Arguments:  []interface{}{"Rafael's", 3, "arguments"},
			Want:       `{"level":"panic","what":"A Rafael's message with 3 arguments.","when":""}`,
			ShouldFail: false,
		},
		{
			Level:      InfoLevel,
			Message:    "A %s message with %d %s.",
			Arguments:  []interface{}{"Rafael's", "arguments"},
			Want:       `{"level":"info","what":"A Rafael's message with arguments.","when":""}`,
			ShouldFail: true,
		},
	}

	for i, test := range tests {
		f := format(test.Message, test.Arguments...)
		m := serializeToFile(test.Level, f, "")

		if ok := reflect.DeepEqual(test.Want, string(m)); !ok && !test.ShouldFail {
			t.Logf("Failed assertion %d. Wants: %s | Got %s", i, test.Want, m)
		}
	}
}