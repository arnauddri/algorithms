package matrix

import (
	"testing"
)

func TestStrassen(t *testing.T) {
	A := makeMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	B := makeMatrix([]float64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 4, 4)

	C := strassen(A, B)
	D := multiplyMatrix(A, B)
	if !FloatArrayEquals(C.elements, D.elements) {
		t.Error()
	}

	A = makeMatrix([]float64{1, 1, 1, 1}, 2, 2)
	B = makeMatrix([]float64{1, 1, 1, 1}, 2, 2)

	C = strassen(A, B)
	D = multiplyMatrix(A, B)
	if !FloatArrayEquals(C.elements, D.elements) {
		t.Error()
	}

	A = makeMatrix([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3, 3)
	B = makeMatrix([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3, 3)

	C = strassen(A, B)
	D = multiplyMatrix(A, B)
	if !FloatArrayEquals(C.elements, D.elements) {
		t.Error()
	}
}
