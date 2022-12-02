package benchexample

import "testing"

// Wrong benchmarks:

func BenchmarkFibWrong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n) // <- wrong argument
	}
}

func BenchmarkFibWrong2(b *testing.B) {
	Fib(b.N)
}
