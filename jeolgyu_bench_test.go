package jeolgyu

import (
	"os"
	"testing"
)

var (
	temp = os.TempDir()
)

func BenchmarkSinkFile(b *testing.B) {
	l, err := New(Settings{SinkType: SinkFile, Filepath: temp})

	if err != nil {
		b.Fatalf("Couldn't initialize logger. Reason %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info("Such meaningful %d ith message, wow.", i)
	}
}

func BenchmarkSinkOutput(b *testing.B) {
	l, err := New(Settings{SinkType: SinkOutput})

	if err != nil {
		b.Fatalf("Couldn't initialize logger. Reason %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info("Such meaningful %d ith message, wow.", i)
	}
}

func BenchmarkSinkBoth(b *testing.B) {
	l, err := New(Settings{SinkType: SinkFile, Filepath: temp})

	if err != nil {
		b.Fatalf("Couldn't initialize logger. Reason %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info("Such meaningful %d ith message, wow.", i)
	}
}
