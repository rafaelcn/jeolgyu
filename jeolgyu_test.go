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
}

func TestSinkFile(t *testing.T) {
	j, err := New(Settings{SinkType: SinkFile, Filepath: ""})

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

		// strip time from the last line
		s := strip(last, t)

		if ok := reflect.DeepEqual(tt.Want, s); !ok {
			t.Errorf("Failed assertion %d. Wants: %s | Got %s", i, tt.Want, s)
		}
	}
}

func TestFilename(t *testing.T) {
	f := "something"
	p := "/tmp/"
	l, err := New(Settings{SinkType: SinkFile, Filepath: p, Filename: f})

	if err != nil {
		t.Error(err)
	}

	_, err = ioutil.ReadFile(l.filename)

	if err != nil {
		t.Errorf("Failed to read file %s. Reason %v", f, err)
	}
}

func TestFilenameWrite(t *testing.T) {
	f := "something"
	p := "/tmp/"
	l, err := New(Settings{SinkType: SinkFile, Filepath: p, Filename: f})

	if err != nil {
		t.Error(err)
	}

	for i, tt := range jeolgyuTests {

		switch tt.Level {
		case InfoLevel:
			l.Info(tt.Message, tt.Arguments...)
		case WarningLevel:
			l.Warning(tt.Message, tt.Arguments...)
		case ErrorLevel:
			l.Error(tt.Message, tt.Arguments...)
		case PanicLevel:
			l.Panic(tt.Message, tt.Arguments...)
		}

		content, err := ioutil.ReadFile(l.filename)

		if err != nil {
			t.Fatalf("Log file couldn't be opened. Reason %v", err)
		}

		lines := strings.Split(string(content), "\n")
		last := lines[len(lines)-2]

		// strip time from the last line
		s := strip(last, t)

		if ok := reflect.DeepEqual(tt.Want, s); !ok {
			t.Errorf("Failed assertion %d. Wants: %s | Got %s", i, tt.Want, s)
		}
	}
}

// strip removes the timestamp of a given string
func strip(m string, t *testing.T) string {
	var b strings.Builder
	b.Grow(len(m))

	count := 0

	for _, v := range m {
		// corresponding value to a quote character (")
		if v == 34 {
			count++
		}

		if count != 11 {
			_, err := b.WriteRune(v)

			if err != nil {
				t.Errorf("Failed to write rune %v. Reason %v", v, err)
			}
		}
	}

	s := b.String()
	s = s[:b.Len()-1] + `"}`

	return s
}
