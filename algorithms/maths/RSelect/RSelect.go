package rselect

import (
	"math/rand"
)

// RSelect returns the i-th smallest (0-indexed) of the first n elements of arr.
// arr is reordered in place.
func RSelect(arr []int, n, i int) int {
	if n == 1 {
		return arr[0]
	}

	pivot := rand.Intn(n)
	v := arr[pivot]
	arr[pivot], arr[n-1] = arr[n-1], arr[pivot]

	left := 0
	for j := 0; j < n-1; j++ {
		if arr[j] < v {
			arr[j], arr[left] = arr[left], arr[j]
			left++
		}
	}
	arr[left], arr[n-1] = arr[n-1], arr[left]

	switch {
	case left == i:
		return arr[left]
	case left > i:
		return RSelect(arr[:left], left, i)
	default:
		return RSelect(arr[left+1:n], n-left-1, i-left-1)
	}
}
