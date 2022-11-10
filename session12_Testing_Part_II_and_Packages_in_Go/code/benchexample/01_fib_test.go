package benchexample

import "testing"

// RUN from shell: go test -bench=.

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
