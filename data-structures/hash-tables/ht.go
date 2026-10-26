// Implementation of a Hash Table with Separate Chaining
// with linked lists using Horner's hash function
// http://en.wikipedia.org/wiki/Hash_table#Separate_chaining_with_linked_lists
package ht

import (
	"errors"
	"github.com/arnauddri/algorithms/data-structures/linked-list"
	"math"
)

type HashTable struct {
	Table    map[int]*list.List
	Size     int
	Capacity int
}

type item struct {
	key   string
	value interface{}
}

func New(cap int) *HashTable {
	table := make(map[int]*list.List, cap)
	return &HashTable{Table: table, Size: 0, Capacity: cap}
}

func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.position(key)
	item, err := ht.find(index, key)

	if item == nil {
		return "", errors.New("Not Found")
	}

	return item.value, err
}

func (ht *HashTable) Put(key, value string) {
	ht.put(key, value)
}

func (ht *HashTable) put(key string, value interface{}) {
	index := ht.position(key)

	if ht.Table[index] == nil {
		ht.Table[index] = list.NewList()
	}

	item := &item{key: key, value: value}

	a, err := ht.find(index, key)
	if err != nil {
		// The key doesn't exist in HashTable
		ht.Table[index].Append(item)
		ht.Size++
	} else {
		// The key exists so we overwrite its value
		a.value = value
	}
}

func (ht *HashTable) Del(key string) error {
	index := ht.position(key)
	l := ht.Table[index]
	var val *item

	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	if val == nil {
		return nil
	}

	ht.Size--
	return l.Remove(val)
}

func (ht *HashTable) ForEach(f func(*item)) {
	for k := range ht.Table {
		if ht.Table[k] != nil {
			ht.Table[k].Each(func(node list.Node) {
				f(node.Value.(*item))
			})
		}
	}
}

func (ht *HashTable) position(s string) int {
	return hashCode(s) % ht.Capacity
}

func (ht *HashTable) find(i int, key string) (*item, error) {
	l := ht.Table[i]
	var val *item

	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	if val == nil {
		return nil, errors.New("Not Found")
	}

	return val, nil
}

// Horner's Method to hash string of length L (O(L))
func hashCode(s string) int {
	hash := int32(0)
	for i := 0; i < len(s); i++ {
		hash = int32(hash<<5-hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}

// Resize rehashes every item into a table of capacity newCap
func (ht *HashTable) Resize(newCap int) error {
	if newCap <= 0 {
		return errors.New("Capacity must be positive")
	}
	old := ht.Table
	ht.Table = make(map[int]*list.List, newCap)
	ht.Capacity = newCap
	ht.Size = 0
	for _, l := range old {
		l.Each(func(node list.Node) {
			it := node.Value.(*item)
			ht.put(it.key, it.value)
		})
	}
	return nil
}
