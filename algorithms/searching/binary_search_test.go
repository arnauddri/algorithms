package bs

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	sorted := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		sorted[i] = 2 * i
	}

	for i := 0; i < 10000; i++ {
		index := search(sorted, 2*i)

		if index != i {
			fmt.Println(index)
			t.Error()
		}
	}

	if search(sorted, 3) != -1 {
		t.Error()
	}
}
