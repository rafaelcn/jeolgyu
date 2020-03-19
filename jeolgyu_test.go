package jeolgyu

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

type Test struct {
	Level     Level
	Message   string
	Arguments []interface{}
	Want      string
}

var jeolgyuTests []Test = []Test{
	{
		Level:   PanicLevel,
		Message: "Something very bad happened",
		Want:    `{"level":"panic","what":"Something very bad happened",when:""}`,
	},
	{
		Level:   WarningLevel,
		Message: "Something bad might have happened",
		Want:    `{"level":"warning","what":"Something bad might have happened",when:""}`,
	},
	{
		Level:   ErrorLevel,
		Message: "ERRRRRRRRROER o.O",
		Want:    `{"level":"error","what":"ERRRRRRRRROER o.O",when:""}`,
	},
	{
		Level:   InfoLevel,
		Message: "Don't worry, everything is fine",
		Want:    `{"level":"info","what":"Don't worry, everything is fine",when:""}`,
	},
	{
		Level:   InfoLevel,
		Message: "Some information for you sir",
		Want:    `{"level":"info","what":"Some information for you sir",when:""}`,
	},
}

func TestSinkFile(t *testing.T) {
	j, err := New(SinkFile, "")

	if err != nil {
		t.Error(err)
	}

	for i, tt := range jeolgyuTests {
		switch tt.Level {
		case InfoLevel:
			j.Info(tt.Message, tt.Arguments...)
		case WarningLevel:
			j.Warning(tt.Message, tt.Arguments...)
		case ErrorLevel:
			j.Error(tt.Message, tt.Arguments...)
		case PanicLevel:
			j.Panic(tt.Message, tt.Arguments...)
		}

		content, err := ioutil.ReadFile(j.filename)

		if err != nil {
			t.Fatalf("Log file couldn't be opened. Reason %v", err)
		}

		lines := strings.Split(string(content), "\n")
		last := lines[len(lines)-2]

		if ok := reflect.DeepEqual(tt.Want, last); !ok {
			t.Logf("Failed assertion %d. Wants: %s | Got %s", i, tt.Want, last)
		}
	}
}
