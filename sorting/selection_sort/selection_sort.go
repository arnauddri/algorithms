package main

import (
	"fmt"
)

func main() {
	list := []int{1, 3, 2, 8, 4, 10, 54, 49, 6}
	fmt.Println("unsorted: ", list)

	selection_sort(list)
	fmt.Println("sorted: ", list)
}

func selection_sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
}
