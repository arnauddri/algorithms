package permutations

import ()

func recursiveCount(a []int) ([]int, int) {
	n := len(a)
	if n < 2 {
		return a, 0
	}

	b, left := recursiveCount(a[:n>>1])
	c, right := recursiveCount(a[n>>1:])
	d := make([]int, 0)

	i, j := 0, 0
	inversions := 0

	for k := 0; k < n; k++ {
		if b[i] < c[j] {
			d = append(d, b[i])
			i++
			if i == len(b) {
				for j < len(c) {
					d = append(d, c[j])
					j++
				}
				break
			}
		} else {
			d = append(d, c[j])
			j++
			inversions += len(b) - i
			if j == len(c) {
				for i < len(b) {
					d = append(d, b[i])
					i++
				}
				break
			}
		}
	}

	return d, left + right + inversions
}

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
