package matrix

import (
	"testing"
)

func TestMakeMatrix(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5, 6}
	A := MakeMatrix(a, 3, 2)

	if A.cols != 2 ||
		A.rows != 3 ||
		!FloatArrayEquals(A.Elements, a) {
		t.Error()
	}
}

func TestCount(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5, 6}
	A := MakeMatrix(a, 3, 2)

	if A.CountRows() != 3 {
		t.Error()
	}
	if A.CountCols() != 2 {
		t.Error()
	}
}

func TestGetElm(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5, 6}
	A := MakeMatrix(a, 3, 2)

	for i := 0; i < 3; i++ {
		if A.GetElm(0, i) != float64(i+1) {
			t.Error()
		}
	}
}

func TestSetElm(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5, 6}
	A := MakeMatrix(a, 3, 2)

	A.SetElm(0, 0, 10)

	if A.GetElm(0, 0) != 10 {
		t.Error()
	}
}

func TestTrace(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	A := MakeMatrix(a, 2, 2)

	if A.trace() != 5 {
		t.Error()
	}
}

func TestAdd(t *testing.T) {
	a := []float64{1, 1, 1, 1}
	A := MakeMatrix(a, 2, 2)
	B := MakeMatrix([]float64{2, 2, 2, 2}, 2, 2)

	A.add(A)
	if !FloatArrayEquals(A.Elements, B.Elements) {
		t.Error()
	}
}

func TestSubstract(t *testing.T) {
	a := []float64{1, 1, 1, 1}
	A := MakeMatrix(a, 2, 2)
	B := MakeMatrix([]float64{2, 2, 2, 2}, 2, 2)

	B.substract(A)
	if !FloatArrayEquals(A.Elements, B.Elements) {
		t.Error()
	}
}

func TestScale(t *testing.T) {
	a := []float64{1, 1, 1, 1}
	A := MakeMatrix(a, 2, 2)
	B := MakeMatrix([]float64{2, 2, 2, 2}, 2, 2)

	A.scale(2)

	if !FloatArrayEquals(A.Elements, B.Elements) {
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
