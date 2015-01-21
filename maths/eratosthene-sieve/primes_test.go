package primes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAllPrimesTo(t *testing.T) {
	primes := getAllPrimesTo(60)
	expected := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	if !reflect.DeepEqual(primes, expected) {
		fmt.Println(primes)
		t.Error()
	}
}

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
