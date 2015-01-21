package permutations

import ()

func iterativeCount(array []int) int {
	n := len(array)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if array[j] < array[i] {
				count++
			}
		}
	}
	return count
}
