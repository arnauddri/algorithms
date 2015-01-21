package gcd

import (
	"testing"
)

func TestRecurse(t *testing.T) {
	if recurse(14, 7) != 7 ||
		recurse(4, 2) != 2 ||
		recurse(31, 2) != 1 ||
		recurse(33, 11) != 11 {
		t.Error()
	}
}

func TestIterative(t *testing.T) {
	if iter(14, 7) != 7 ||
		iter(4, 2) != 2 ||
		iter(31, 2) != 1 ||
		iter(33, 11) != 11 {
		t.Error()
	}
}

func BenchmarkRecurse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurse(131313131, 12343545)
	}
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter(131313131, 12343545)
	}
}
