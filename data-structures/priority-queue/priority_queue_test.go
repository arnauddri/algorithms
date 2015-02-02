package pq

import (
	"fmt"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	h := NewMax()

	h.Insert(*NewItem(8, 10))
	h.Insert(*NewItem(7, 11))
	h.Insert(*NewItem(6, 12))
	h.Insert(*NewItem(3, 13))
	h.Insert(*NewItem(1, 14))
	h.Insert(*NewItem(0, 15))
	h.Insert(*NewItem(2, 16))
	h.Insert(*NewItem(4, 17))
	h.Insert(*NewItem(9, 18))
	h.Insert(*NewItem(5, 19))

	sorted := make([]Item, 0)
	for h.Len() > 0 {
		sorted = append(sorted, h.Extract())
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority < sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestMinPriorityQueue(t *testing.T) {
	h := NewMin()

	h.Insert(*NewItem(8, 10))
	h.Insert(*NewItem(7, 11))
	h.Insert(*NewItem(6, 12))
	h.Insert(*NewItem(3, 13))
	h.Insert(*NewItem(1, 14))
	h.Insert(*NewItem(0, 15))
	h.Insert(*NewItem(2, 16))
	h.Insert(*NewItem(4, 17))
	h.Insert(*NewItem(9, 18))
	h.Insert(*NewItem(5, 19))

	sorted := make([]Item, 0)
	for h.Len() > 0 {
		sorted = append(sorted, h.Extract())
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority > sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestChangePriority(t *testing.T) {
	h := NewMax()

	h.Insert(*NewItem(8, 10))
	h.Insert(*NewItem(7, 11))
	h.Insert(*NewItem(6, 12))
	h.Insert(*NewItem(3, 13))
	h.Insert(*NewItem(1, 14))
	h.Insert(*NewItem(0, 15))
	h.Insert(*NewItem(2, 16))
	h.Insert(*NewItem(4, 17))
	h.Insert(*NewItem(9, 18))
	h.Insert(*NewItem(5, 19))

	h.ChangePriority(8, 66)
	popped := h.Extract()

	if popped.Value != 8 {
		t.Error()
	}
}
