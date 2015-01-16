package matrix

import (
	"math"
)

func strassen(A *matrix, B *matrix) *matrix {
	n := A.countRows()
	bigN := scaleSize(n)

	bigA := makeMatrix(make([]float64, bigN*bigN), bigN, bigN)
	bigB := makeMatrix(make([]float64, bigN*bigN), bigN, bigN)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bigA.setElm(i, j, A.getElm(i, j))
			bigB.setElm(i, j, B.getElm(i, j))
		}
	}

	bigC := recurseStrassen(bigA, bigB)

	C := makeMatrix(make([]float64, n*n), n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C.setElm(i, j, bigC.getElm(i, j))
		}
	}

	return C
}

func recurseStrassen(A *matrix, B *matrix) *matrix {
	n := A.countRows()

	newSize := n / 2

	if n < 2 {
		return makeMatrix([]float64{A.getElm(0, 0) * B.getElm(0, 0)}, 1, 1)
	}

	A11 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A12 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A21 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A22 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	B11 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B12 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B21 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B22 := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			A11.setElm(i, j, A.getElm(i, j))
			A12.setElm(i, j, A.getElm(i, j+newSize))
			A21.setElm(i, j, A.getElm(i+newSize, j))
			A22.setElm(i, j, A.getElm(i+newSize, j+newSize))

			B11.setElm(i, j, B.getElm(i, j))
			B12.setElm(i, j, B.getElm(i, j+newSize))
			B21.setElm(i, j, B.getElm(i+newSize, j))
			B22.setElm(i, j, B.getElm(i+newSize, j+newSize))
		}
	}

	a := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	b := makeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	// P1 = (A11 + A22) * (B11 + B22)
	a = addMatrix(A11, A22)
	b = addMatrix(B11, B22)
	P1 := recurseStrassen(a, b)

	// P2 = (A21 + A22) * (B11)
	a = addMatrix(A21, A22)
	P2 := recurseStrassen(a, B11)

	// P3 = (A11) * (B12 - B11)
	b = substractMatrix(B12, B22)
	P3 := recurseStrassen(A11, b)

	// P4 = A22 * (B21 - B22)
	b = substractMatrix(B21, B11)
	P4 := recurseStrassen(A22, b)

	// P5 = (A11 + A12) * B22
	a = addMatrix(A11, A12)
	P5 := recurseStrassen(a, B22)

	// P6 = (A21 - A11) * (B11 + B12)
	a = substractMatrix(A21, A11)
	b = addMatrix(B11, B12)
	P6 := recurseStrassen(a, b)

	// P7 = (A12 - A22) * (B21 + B22)
	a = substractMatrix(A12, A22)
	b = addMatrix(B21, B22)
	P7 := recurseStrassen(a, b)

	// Aggregates the result into C
	C12 := addMatrix(P3, P5)
	C21 := addMatrix(P2, P4)

	a = addMatrix(P1, P4)
	b = addMatrix(a, P7)
	C11 := substractMatrix(b, P5)

	a = addMatrix(P1, P3)
	b = addMatrix(a, P6)
	C22 := substractMatrix(b, P2)

	C := makeMatrix(make([]float64, n*n), n, n)

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			C.setElm(i, j, C11.getElm(i, j))
			C.setElm(i, j+newSize, C12.getElm(i, j))
			C.setElm(i+newSize, j, C21.getElm(i, j))
			C.setElm(i+newSize, j+newSize, C22.getElm(i, j))
		}
	}

	return C
}

func scaleSize(n int) int {
	log2 := math.Ceil(math.Log(float64(n)) / math.Log(float64(2)))
	return int(math.Pow(2, log2))
}
