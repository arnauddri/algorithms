package strassen

import (
	"github.com/arnauddri/algorithms/data-structures/matrix"
	"math"
)

func Multiply(A *matrix.Matrix, B *matrix.Matrix) *matrix.Matrix {
	n := A.CountRows()
	bigN := scaleSize(n)

	bigA := matrix.MakeMatrix(make([]float64, bigN*bigN), bigN, bigN)
	bigB := matrix.MakeMatrix(make([]float64, bigN*bigN), bigN, bigN)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bigA.SetElm(i, j, A.GetElm(i, j))
			bigB.SetElm(i, j, B.GetElm(i, j))
		}
	}

	bigC := recurse(bigA, bigB)

	C := matrix.MakeMatrix(make([]float64, n*n), n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C.SetElm(i, j, bigC.GetElm(i, j))
		}
	}

	return C
}

func recurse(A *matrix.Matrix, B *matrix.Matrix) *matrix.Matrix {
	n := A.CountRows()

	newSize := n / 2

	if n < 2 {
		return matrix.MakeMatrix([]float64{A.GetElm(0, 0) * B.GetElm(0, 0)}, 1, 1)
	}

	A11 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A12 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A21 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A22 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	B11 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B12 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B21 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B22 := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			A11.SetElm(i, j, A.GetElm(i, j))
			A12.SetElm(i, j, A.GetElm(i, j+newSize))
			A21.SetElm(i, j, A.GetElm(i+newSize, j))
			A22.SetElm(i, j, A.GetElm(i+newSize, j+newSize))

			B11.SetElm(i, j, B.GetElm(i, j))
			B12.SetElm(i, j, B.GetElm(i, j+newSize))
			B21.SetElm(i, j, B.GetElm(i+newSize, j))
			B22.SetElm(i, j, B.GetElm(i+newSize, j+newSize))
		}
	}

	a := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	b := matrix.MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	// P1 = (A11 + A22) * (B11 + B22)
	a = matrix.Add(A11, A22)
	b = matrix.Add(B11, B22)
	P1 := recurse(a, b)

	// P2 = (A21 + A22) * (B11)
	a = matrix.Add(A21, A22)
	P2 := recurse(a, B11)

	// P3 = (A11) * (B12 - B11)
	b = matrix.Substract(B12, B22)
	P3 := recurse(A11, b)

	// P4 = A22 * (B21 - B22)
	b = matrix.Substract(B21, B11)
	P4 := recurse(A22, b)

	// P5 = (A11 + A12) * B22
	a = matrix.Add(A11, A12)
	P5 := recurse(a, B22)

	// P6 = (A21 - A11) * (B11 + B12)
	a = matrix.Substract(A21, A11)
	b = matrix.Add(B11, B12)
	P6 := recurse(a, b)

	// P7 = (A12 - A22) * (B21 + B22)
	a = matrix.Substract(A12, A22)
	b = matrix.Add(B21, B22)
	P7 := recurse(a, b)

	// Aggregates the result into C
	C12 := matrix.Add(P3, P5)
	C21 := matrix.Add(P2, P4)

	a = matrix.Add(P1, P4)
	b = matrix.Add(a, P7)
	C11 := matrix.Substract(b, P5)

	a = matrix.Add(P1, P3)
	b = matrix.Add(a, P6)
	C22 := matrix.Substract(b, P2)

	C := matrix.MakeMatrix(make([]float64, n*n), n, n)

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			C.SetElm(i, j, C11.GetElm(i, j))
			C.SetElm(i, j+newSize, C12.GetElm(i, j))
			C.SetElm(i+newSize, j, C21.GetElm(i, j))
			C.SetElm(i+newSize, j+newSize, C22.GetElm(i, j))
		}
	}

	return C
}

func scaleSize(n int) int {
	log2 := math.Ceil(math.Log(float64(n)) / math.Log(float64(2)))
	return int(math.Pow(2, log2))
}
