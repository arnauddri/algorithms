package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMin()

	h.Insert(Int(8))
	h.Insert(Int(7))
	h.Insert(Int(6))
	h.Insert(Int(3))
	h.Insert(Int(1))
	h.Insert(Int(0))
	h.Insert(Int(2))
	h.Insert(Int(4))
	h.Insert(Int(9))
	h.Insert(Int(5))

	sorted := make([]Int, 0)
	for h.Len() > 0 {
		sorted = append(sorted, h.Extract().(Int))
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] > sorted[i+1] {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestMaxHeap(t *testing.T) {
	h := NewMax()

	h.Insert(Int(8))
	h.Insert(Int(7))
	h.Insert(Int(6))
	h.Insert(Int(3))
	h.Insert(Int(1))
	h.Insert(Int(0))
	h.Insert(Int(2))
	h.Insert(Int(4))
	h.Insert(Int(9))
	h.Insert(Int(5))

	sorted := make([]Int, 0)
	for h.Len() > 0 {
		sorted = append(sorted, h.Extract().(Int))
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] < sorted[i+1] {
			fmt.Println(sorted)
			t.Error()
		}
	}
}
