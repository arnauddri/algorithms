package heap

import (
	"github.com/arnauddri/algorithms/data-structures/heap"
)

func sort(a []int) []int {
	h := heap.NewMin()
	for i := 0; i < len(a); i++ {
		h.Insert(heap.Int(a[i]))
	}

	sorted := make([]int, 0)

	for i := 0; i < len(a); i++ {
		sorted = append(sorted, int(h.Extract().(heap.Int)))
	}

	return sorted
}
