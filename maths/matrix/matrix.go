// Package matrix provides some methods to use with matrix
package matrix

import (
	"errors"
)

type Matrix interface {
	countRows() int
	countCols() int
	getElm(i int, j int) float64
	trace() float64

	setElm(i int, j int, v float64)
	add(*matrix) error
	substract(*matrix) error
	scale(float64)

	copy() []float64
	diagonalCopy() []float64
}

type matrix struct {
	// Number of rows
	rows int
	// Number of columns
	cols int
	// Matrix stored as a flat array: Aij = elements[i*step + j]
	elements []float64
	// Offset between rows
	step int
}

func makeMatrix(elements []float64, rows, cols int) *matrix {
	A := new(matrix)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.elements = elements

	return A
}
func (A *matrix) countRows() int {
	return A.rows
}

func (A *matrix) countCols() int {
	return A.cols
}

func (A *matrix) getElm(i int, j int) float64 {
	return A.elements[i*A.step+j]
}

func (A *matrix) setElm(i int, j int, v float64) {
	A.elements[i*A.step+j] = v
}

func (A *matrix) diagonalCopy() []float64 {
	diag := make([]float64, A.cols)
	for i := 0; i < len(diag); i++ {
		diag[i] = A.getElm(i, i)
	}
	return diag
}

func (A *matrix) copy() *matrix {
	B := new(matrix)
	B.rows = A.rows
	B.cols = A.cols
	B.step = A.step

	B.elements = make([]float64, A.cols*A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.elements[i*A.step+j] = A.getElm(i, j)
		}
	}
	return B
}

func (A *matrix) trace() float64 {
	var tr float64 = 0
	for i := 0; i < A.cols; i++ {
		tr += A.getElm(i, i)
	}
	return tr
}

func (A *matrix) add(B *matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("Wrong input sizes")
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.setElm(i, j, A.getElm(i, j)+B.getElm(i, j))
		}
	}

	return nil
}

func (A *matrix) substract(B *matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("Wrong input sizes")
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.setElm(i, j, A.getElm(i, j)-B.getElm(i, j))
		}
	}

	return nil
}

func (A *matrix) scale(a float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.setElm(i, j, a*A.getElm(i, j))
		}
	}
}

func addMatrix(A *matrix, B *matrix) *matrix {
	result := makeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.setElm(i, j, A.getElm(i, j)+B.getElm(i, j))
		}
	}

	return result
}

func substractMatrix(A *matrix, B *matrix) *matrix {
	result := makeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.setElm(i, j, A.getElm(i, j)-B.getElm(i, j))
		}
	}

	return result
}

func multiplyMatrix(A *matrix, B *matrix) *matrix {
	result := makeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.getElm(i, k) * B.getElm(k, j)
			}
			result.setElm(i, j, sum)
		}
	}

	return result
}
