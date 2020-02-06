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
	tt := []serializerTests{
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
			Arguments:  []interface{}{"Rafael's", 3, "arguments"},
			Want:       `{"level":"info","what":"A Rafael's message with arguments.","when":""}`,
			ShouldFail: true,
		},
	}

	for i, te := range tt {
		f := format(te.Message, te.Arguments...)
		m := Serialize(te.Level, f, "")

		if ok := reflect.DeepEqual(te.Want, []byte(m)); ok && !te.ShouldFail {
			t.Logf("Failed assertion %d. Wants: %s | Got %s", i, te.Want, m)
		}
	}
}
