package insertion

import ()

func sort(a []int) {
	for i := 1; i < len(a); i++ {
		value := a[i]
		j := i - 1
		for j >= 0 && a[j] > value {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = value
	}
}
