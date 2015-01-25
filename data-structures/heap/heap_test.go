package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := New()

	h.Insert(8)
	h.Insert(7)
	h.Insert(6)
	h.Insert(3)
	h.Insert(1)
	h.Insert(0)
	h.Insert(2)
	h.Insert(4)
	h.Insert(9)
	h.Insert(5)
	fmt.Println(h.data)
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
}
