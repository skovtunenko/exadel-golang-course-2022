package benchexample

import "testing"

// Various inputs:

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

// Typical output:

// BenchmarkFib1   1000000000               2.84 ns/op
// BenchmarkFib2   500000000                7.92 ns/op
// BenchmarkFib3   100000000               13.0 ns/op
// BenchmarkFib10   5000000               447 ns/op
// BenchmarkFib20     50000             55668 ns/op
// BenchmarkFib40         2         942888676 ns/op
