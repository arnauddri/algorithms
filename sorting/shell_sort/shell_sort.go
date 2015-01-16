// Package main provides ...
package main

import (
	"fmt"
)

var arr = []int{20, 43, 52, -1, 43, 29, 34}

func main() {
	fmt.Println("Unsorted: ", arr)
	shell_sort(arr)
	fmt.Println("Sorted: ", arr)
}

func shell_sort(arr []int) {
	increment := len(arr) / 2
	for increment > 0 {
		for i := increment; i < len(arr); i++ {
			j := i
			temp := arr[i]

			for j >= increment && arr[j-increment] > temp {
				arr[j] = arr[j-increment]
				j = j - increment
			}
			arr[j] = temp
		}
		if increment == 2 {
			increment = 1
		} else {
			increment = int(increment * 5 / 11)
		}
	}
}
