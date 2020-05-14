package jeolgyu

import (
	"io/ioutil"
	"os"
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

var jeolgyuSinkFileTests []Test = []Test{
	{
		Level:   PanicLevel,
		Message: "Something very bad happened",
		Want:    `{"level":"panic","what":"Something very bad happened","when":""}`,
	},
	{
		Level:   WarningLevel,
		Message: "Something bad might have happened",
		Want:    `{"level":"warning","what":"Something bad might have happened","when":""}`,
	},
	{
		Level:   ErrorLevel,
		Message: "ERRRRRRRRROER o.O",
		Want:    `{"level":"error","what":"ERRRRRRRRROER o.O","when":""}`,
	},
	{
		Level:   InfoLevel,
		Message: "Don't worry, everything is fine",
		Want:    `{"level":"info","what":"Don't worry, everything is fine","when":""}`,
	},
	{
		Level:   InfoLevel,
		Message: "Some information for you sir",
		Want:    `{"level":"info","what":"Some information for you sir","when":""}`,
	},
	{
		Level:     InfoLevel,
		Message:   "Take a look",
		Arguments: []interface{}{"unused argument in the string"},
		Want:      `{"level":"info","what":"Take a look%!(EXTRA string=unused argument in the string)","when":""}`,
	},
	{
		Level:     ErrorLevel,
		Message:   "The author %s is a very good friend",
		Arguments: []interface{}{"Mikael Messias"},
		Want:      `{"level":"error","what":"The author Mikael Messias is a very good friend","when":""}`,
	},
}

var jeolgyuSinkOutputTests []Test = []Test{
	{
		Level:   PanicLevel,
		Message: "Something very bad happened",
		Want:    `[P] Something very bad happened"`,
	},
	{
		Level:   WarningLevel,
		Message: "Something bad might have happened",
		Want:    `[!] Something bad might have happened`,
	},
	{
		Level:   ErrorLevel,
		Message: "ERRRRRRRRROER o.O",
		Want:    `[x] ERRRRRRRRROER o.O`,
	},
	{
		Level:   InfoLevel,
		Message: "Don't worry, everything is fine",
		Want:    `[+] Don't worry, everything is fine`,
	},
	{
		Level:   InfoLevel,
		Message: "Some information for you sir",
		Want:    `[+] Some information for you sir`,
	},
}

func TestSinkFile(t *testing.T) {
	j, err := New(SinkFile, "")
	j.testing = true

	if err != nil {
		t.Error(err)
	}

	for i, tt := range jeolgyuSinkFileTests {
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

func testSinkOutput(t *testing.T) {
	// FIXME: Mock an stdout to use as testing because in the current status
	// there's no way to run this test.

	j, err := New(SinkOutput, "")
	j.testing = true

	if err != nil {
		t.Error(err)
	}

	for i, tt := range jeolgyuSinkOutputTests {
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

		buf := make([]byte, len(tt.Message))

		_, err := os.Stdout.Read(buf)

		if err != nil {
			t.Errorf("Couldn't read from stdout. Reason %v", err)
		}

		output := string(buf)

		if ok := reflect.DeepEqual(tt.Want, output); !ok {
			t.Logf("Failed assertion %d. Wants: %s | Got %s", i, tt.Want, output)
		}
	}
}
