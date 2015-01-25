package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	list := []int{1, 3, 2, 8, 4, 10, 54, 49, 6}

	sorted := sort(list)
	sorted_list := []int{1, 2, 3, 4, 6, 8, 10, 49, 54}

	if !deepEqual(sorted, sorted_list) {
		fmt.Println(sorted)
		t.Error()
	}
}

func deepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
