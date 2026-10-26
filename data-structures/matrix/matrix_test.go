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

func TestAddWrongSizes(t *testing.T) {
	A := MakeMatrix([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	B := MakeMatrix([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	C := MakeMatrix([]float64{1, 2, 3}, 1, 3)
	if A.add(C) == nil || A.substract(C) == nil {
		t.Error("expected error for mismatched rows")
	}
	if A.add(B) == nil || A.substract(B) == nil {
		t.Error("expected error for mismatched shapes")
	}
}

func TestMultiply(t *testing.T) {
	A := MakeMatrix([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	B := MakeMatrix([]float64{7, 8, 9, 10, 11, 12}, 3, 2)
	C := Multiply(A, B)
	if C.CountRows() != 2 || C.CountCols() != 2 ||
		!FloatArrayEquals(C.Elements, []float64{58, 64, 139, 154}) {
		t.Errorf("Multiply() = %v (%dx%d)", C.Elements, C.CountRows(), C.CountCols())
	}
}

func TestTranspose(t *testing.T) {
	A := MakeMatrix([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	B := Transpose(A)
	if B.CountRows() != 3 || B.CountCols() != 2 ||
		!FloatArrayEquals(B.Elements, []float64{1, 4, 2, 5, 3, 6}) {
		t.Errorf("Transpose() = %v", B.Elements)
	}
}
