package strassen

import (
	"github.com/arnauddri/algorithms/data-structures/matrix"
	"testing"
)

func TestStrassen(t *testing.T) {
	A := matrix.MakeMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	B := matrix.MakeMatrix([]float64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 4, 4)

	C := Multiply(A, B)
	D := matrix.Multiply(A, B)
	if !FloatArrayEquals(C.Elements, D.Elements) {
		t.Error()
	}

	A = matrix.MakeMatrix([]float64{1, 1, 1, 1}, 2, 2)
	B = matrix.MakeMatrix([]float64{1, 1, 1, 1}, 2, 2)

	C = Multiply(A, B)
	D = matrix.Multiply(A, B)
	if !FloatArrayEquals(C.Elements, D.Elements) {
		t.Error()
	}

	A = matrix.MakeMatrix([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3, 3)
	B = matrix.MakeMatrix([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3, 3)

	C = Multiply(A, B)
	D = matrix.Multiply(A, B)
	if !FloatArrayEquals(C.Elements, D.Elements) {
		t.Error()
	}
}

func FloatArrayEquals(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkStrassen(b *testing.B) {
	a := make([]float64, 100000000, 100000000)
	c := make([]float64, 100000000, 100000000)
	for i := 0; i < 100000000; i++ {
		a[i] = float64(i)
		c[i] = float64(100000000 - i)
	}
	A := matrix.MakeMatrix(a, 10000, 10000)
	B := matrix.MakeMatrix(c, 10000, 10000)

	for i := 0; i < b.N; i++ {
		Multiply(A, B)
	}
}

func BenchmarkMultiply(b *testing.B) {
	a := make([]float64, 100000000, 100000000)
	c := make([]float64, 100000000, 100000000)

	for i := 0; i < 100000000; i++ {
		a[i] = float64(i)
		c[i] = float64(100000000 - i)
	}
	A := matrix.MakeMatrix(a, 10000, 10000)
	B := matrix.MakeMatrix(c, 10000, 10000)

	for i := 0; i < b.N; i++ {
		matrix.Multiply(A, B)
	}
}
