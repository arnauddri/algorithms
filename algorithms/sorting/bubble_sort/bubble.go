// Package main provides ...
package bubble

import (
//"fmt"
)

func sort(a []int) {
	for itemCount := len(a) - 1; ; itemCount-- {
		swap := false
		for i := 1; i <= itemCount; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}
