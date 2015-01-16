package rselect

import (
	"math/rand"
)

func RSelect(arr []int, n, i int) int {
	pivot := rand.Intn(n)
	v := arr[pivot]
	left := 0

	arr[pivot], arr[n-1] = arr[n-1], arr[pivot]

	for j := 0; j < n-1; j++ {
		if arr[j] <= v {
			arr[j], arr[left] = arr[left], arr[j]
			left++
		}
	}

	if left+1 >= n {
		return arr[n-1]
	}

	arr[left+1], arr[n-1] = arr[n-1], arr[left+1]

	j := left

	if j < 2 {
		return arr[j]
	}

	if j == i {
		return arr[j]
	} else if j > i {
		return RSelect(arr[:j], j-1, i)
	} else {
		return RSelect(arr[j:], n-j, i-j)
	}
}
