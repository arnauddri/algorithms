// Package Matrix provides some methods to use with Matrix
package matrix

import (
	"errors"
)

//type Matrix interface {
//CountRows() int
//CountCols() int
//GetElm(i int, j int) float64
//trace() float64

//SetElm(i int, j int, v float64)
//add(*Matrix) error
//substract(*Matrix) error
//scale(float64)

//copy() []float64
//diagonalCopy() []float64
//}

type Matrix struct {
	// Number of rows
	rows int
	// Number of columns
	cols int
	// Matrix stored as a flat array: Aij = Elements[i*step + j]
	Elements []float64
	// Offset between rows
	step int
}

func MakeMatrix(Elements []float64, rows, cols int) *Matrix {
	A := new(Matrix)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.Elements = Elements

	return A
}
func (A *Matrix) CountRows() int {
	return A.rows
}

func (A *Matrix) CountCols() int {
	return A.cols
}

func (A *Matrix) GetElm(i int, j int) float64 {
	return A.Elements[i*A.step+j]
}

func (A *Matrix) SetElm(i int, j int, v float64) {
	A.Elements[i*A.step+j] = v
}

func (A *Matrix) diagonalCopy() []float64 {
	diag := make([]float64, A.cols)
	for i := 0; i < len(diag); i++ {
		diag[i] = A.GetElm(i, i)
	}
	return diag
}

func (A *Matrix) copy() *Matrix {
	B := new(Matrix)
	B.rows = A.rows
	B.cols = A.cols
	B.step = A.step

	B.Elements = make([]float64, A.cols*A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.Elements[i*A.step+j] = A.GetElm(i, j)
		}
	}
	return B
}

func (A *Matrix) trace() float64 {
	var tr float64 = 0
	for i := 0; i < A.cols; i++ {
		tr += A.GetElm(i, i)
	}
	return tr
}

func (A *Matrix) add(B *Matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("Wrong input sizes")
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElm(i, j, A.GetElm(i, j)+B.GetElm(i, j))
		}
	}

	return nil
}

func (A *Matrix) substract(B *Matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("Wrong input sizes")
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElm(i, j, A.GetElm(i, j)-B.GetElm(i, j))
		}
	}

	return nil
}

func (A *Matrix) scale(a float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElm(i, j, a*A.GetElm(i, j))
		}
	}
}

func Add(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.SetElm(i, j, A.GetElm(i, j)+B.GetElm(i, j))
		}
	}

	return result
}

func Substract(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.SetElm(i, j, A.GetElm(i, j)-B.GetElm(i, j))
		}
	}

	return result
}

func Multiply(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.GetElm(i, k) * B.GetElm(k, j)
			}
			result.SetElm(i, j, sum)
		}
	}

	return result
}
