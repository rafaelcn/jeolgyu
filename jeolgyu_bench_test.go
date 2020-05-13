package jeolgyu

import "testing"

const (
	N  = 1024
	N2 = 4096
	N3 = 65536
)

// BenchmarkSinkFileN benchmarks the execution of consecutive log.Info calls
// with size N. The following functions just change the ammount of calls it will
// make to the logger.

func BenchmarkSinkFileN(b *testing.B) {

	l, err := New(SinkFile, "/tmp/")

	if err != nil {
		b.Fatalf("Couldn't initialize logger. Reason %v", err)
	}

	b.ResetTimer()

	for i := 0; i < N; i++ {
		l.Info("Such meaningful %d ith message, wow.", i)
	}
}

func BenchmarkSinkFileN2(b *testing.B) {

	l, err := New(SinkFile, "/tmp/")

	if err != nil {
		b.Fatalf("Couldn't initialize logger. Reason %v", err)
	}

	b.ResetTimer()

	for i := 0; i < N2; i++ {
		l.Info("Such meaningful %d ith message, wow.", i)
	}
}

func BenchmarkSinkFileN3(b *testing.B) {

	l, err := New(SinkFile, "/tmp/")

	if err != nil {
		b.Fatalf("Couldn't initialize logger. Reason %v", err)
	}

	b.ResetTimer()

	for i := 0; i < N3; i++ {
		l.Info("Such meaningful %d ith message, wow.", i)
	}
}
