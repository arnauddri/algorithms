package primes

import (
	"testing"
)

func benchmarkGetAllPrimesTo(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAllPrimesTo(n)
	}
}

func BenchmarkGetAllPrimesTo10(b *testing.B)        { benchmarkGetAllPrimesTo(10, b) }
func BenchmarkGetAllPrimesTo100(b *testing.B)       { benchmarkGetAllPrimesTo(100, b) }
func BenchmarkGetAllPrimesTo1000(b *testing.B)      { benchmarkGetAllPrimesTo(1000, b) }
func BenchmarkGetAllPrimesTo10000(b *testing.B)     { benchmarkGetAllPrimesTo(10000, b) }
func BenchmarkGetAllPrimesTo100000(b *testing.B)    { benchmarkGetAllPrimesTo(100000, b) }
func BenchmarkGetAllPrimesTo1000000(b *testing.B)   { benchmarkGetAllPrimesTo(1000000, b) }
func BenchmarkGetAllPrimesTo10000000(b *testing.B)  { benchmarkGetAllPrimesTo(10000000, b) }
func BenchmarkGetAllPrimesTo100000000(b *testing.B) { benchmarkGetAllPrimesTo(100000000, b) }
