package heap

import (
	//"fmt"
	"sync"
)

type Heap struct {
	sync.Mutex
	data []int
}

func New() *Heap {
	return &Heap{
		data: make([]int, 0),
	}
}

func (h *Heap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Get(n int) int {
	return h.data[n]
}

func (h *Heap) Insert(n int) {
	h.Lock()
	defer h.Unlock()

	h.data = append(h.data, n)
	h.siftUp()

	return
}

func (h *Heap) Extract() (el int) {
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

	h.data = append([]int{last}, h.data[1:h.Len()-1]...)
	h.siftDown()

	return
}

func (h *Heap) siftUp() {
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i >> 1
		if h.Get(parent) > h.Get(i) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	for i, child := 0, 1; i < h.Len() && i<<1+1 < h.Len(); i = child {
		child = i<<1 + 1

		if child+1 <= h.Len()-1 && h.Get(child+1) < h.Get(child) {
			child++
		}

		if h.Get(i) < h.Get(child) {
			break
		}

		h.data[i], h.data[child] = h.data[child], h.data[i]
	}
}
