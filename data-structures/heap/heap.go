package heap

import (
	"sync"
)

type Item interface {
	Less(than Item) bool
}

type Heap struct {
	sync.Mutex
	data []Item
	min  bool
}

func New() *Heap {
	return &Heap{
		data: make([]Item, 0),
	}
}

func NewMin() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  true,
	}
}

func NewMax() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  false,
	}
}

func (h *Heap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Get(n int) Item {
	return h.data[n]
}

func (h *Heap) Insert(n Item) {
	h.Lock()
	defer h.Unlock()

	h.data = append(h.data, n)
	h.siftUp()

	return
}

func (h *Heap) Extract() (el Item) {
	h.Lock()
	defer h.Unlock()
	if h.Len() == 0 {
		return
	}

	el = h.data[0]
	last := h.data[h.Len()-1]
	if h.Len() == 1 {
		h.data = nil
		return
	}

	h.data = append([]Item{last}, h.data[1:h.Len()-1]...)
	h.siftDown()

	return
}

func (h *Heap) siftUp() {
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i >> 1
		if h.Less(h.Get(i), h.Get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	for i, child := 0, 1; i < h.Len() && i<<1+1 < h.Len(); i = child {
		child = i<<1 + 1

		if child+1 <= h.Len()-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child++
		}

		if h.Less(h.Get(i), h.Get(child)) {
			break
		}

		h.data[i], h.data[child] = h.data[child], h.data[i]
	}
}

func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}
