package pq

import (
	"github.com/arnauddri/algorithms/data-structures/heap"
)

type Item struct {
	Value    interface{}
	Priority int
}

func NewItem(value interface{}, priority int) (i *Item) {
	return &Item{
		Value:    value,
		Priority: priority,
	}
}

func (x Item) Less(than heap.Item) bool {
	return x.Priority < than.(Item).Priority
}

type PQ struct {
	data heap.Heap
}

func New() (q *PQ) {
	return &PQ{
		data: *heap.NewMax(),
	}
}

func (pq *PQ) Len() int {
	return pq.data.Len()
}

func (pq *PQ) Insert(el Item) {
	pq.data.Insert(heap.Item(el))
}

func (pq *PQ) Extract() (el Item) {
	return pq.data.Extract().(Item)
}
