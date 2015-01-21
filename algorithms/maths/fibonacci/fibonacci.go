package fibonacci

import (
	"math"
	"strconv"
)

func getIter(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	for i, j, k := 1, 1, 1; k < n; i, j, k = i+j, i, k+1 {
		a = i
	}

	return a
}

func getRecurse(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	return getRecurse(n-1) + getRecurse(n-2)
}

func getMatrix(n int) int {
	memo := make(map[int][4]int)

	return _getMatrix(n, memo)
}

func _getMatrix(n int, memo map[int][4]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	power := string(strconv.FormatInt(int64(n-1), 2))
	powers := make([]int, 0)

	for i, r := range power {
		if string(r) == "1" {
			powers = append([]int{int(math.Pow(2, float64(len(power)-i-1)))}, powers...)
		}
	}

	Q := [4]int{1, 1, 1, 0}
	matrices := make([][4]int, 0)

	for _, p := range powers {
		matrices = append(matrices, _raiseToPower(Q, p, memo))
	}

	for len(matrices) > 1 {
		M1 := matrices[0]
		M2 := matrices[1]
		R := _multiplyMatrix(M1, M2)
		matrices = matrices[2:len(matrices)]
		matrices = append(matrices, R)
	}

	return matrices[0][0]
}

func _multiplyMatrix(A [4]int, B [4]int) [4]int {
	var result [4]int

	result[0] = A[0]*B[0] + A[1]*B[2]
	result[1] = A[0]*B[1] + A[1]*B[3]
	result[2] = A[2]*B[0] + A[3]*B[2]
	result[3] = A[2]*B[1] + A[3]*B[3]

	return result
}

func _raiseToPower(A [4]int, p int, memo map[int][4]int) [4]int {
	var result [4]int

	if p == 1 {
		return A
	}

	if cache, ok := memo[p]; ok {
		return cache
	}

	result = _raiseToPower(A, p/2, memo)
	R := _multiplyMatrix(result, result)

	memo[p] = R

	return R
}
