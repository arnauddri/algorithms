package main

import (
	"fmt"
)

func main() {
	list := []int{1, 3, 2, 8, 4, 10, 54, 49, 6}
	fmt.Println("unsorted: ", list)

	insertion_sort(list)
	fmt.Println("sorted: ", list)
}

func insertion_sort(a []int) {
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
