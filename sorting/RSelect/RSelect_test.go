package rselect

import (
	"fmt"
	"testing"
)

func TestRSelect(t *testing.T) {
	arr := []int{5, 2, 6, 8, 3, 1}

	i0 := RSelect(arr, 6, 0)
	i1 := RSelect(arr, 6, 1)
	i2 := RSelect(arr, 6, 2)
	i3 := RSelect(arr, 6, 3)
	i4 := RSelect(arr, 6, 4)
	i5 := RSelect(arr, 6, 5)
	if i0 != 1 &&
		i1 != 2 &&
		i2 != 3 &&
		i3 != 5 &&
		i4 != 6 &&
		i5 != 8 {
		fmt.Println("Error: ", i0, i1, i2, i3, i4, i5)
		t.Error()
	}
}
