package list

import (
//"fmt"
)

type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

func NewList() *List {
	l := new(List)
	l.Length = 0
	return l
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Length == 0
}

func (l *List) Prepend(value interface{}) {
	node := NewNode(value)
	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerHead := l.Head
		formerHead.Prev = node

		node.Next = formerHead
		l.Head = node
	}

	l.Length++
}

func (l *List) Append(value interface{}) {
	node := NewNode(value)

	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerTail := l.Tail
		formerTail.Next = node

		node.Prev = formerTail
		l.Tail = node
	}

	l.Length++
}

func (l *List) Add(value interface{}, index int) {
	if index > l.Len() {
		panic("Index out of range")
	}

	node := NewNode(value)

	if l.Len() == 0 || index == 0 {
		l.Prepend(value)
		return
	}

	if l.Len()-1 == index {
		l.Append(value)
		return
	}

	nextNode := l.Get(index)
	prevNode := nextNode.Prev

	prevNode.Next = node
	node.Prev = prevNode

	nextNode.Prev = node
	node.Next = nextNode

	l.Length++
}

func (l *List) Remove(node *Node) {
	if l.Len() == 0 {
		panic("Empty list")
	}

}

func (l *List) Get(index int) *Node {
	if index > l.Len() {
		panic("Index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

func (l *List) Find(node *Node) int {
	if l.Len() == 0 {
		panic("Empty list")
	}

	index := 0
	found := -1
	l.Map(func(n *Node) {
		index++
		if n.Value == node.Value && found == -1 {
			found = index
		}
	})

	if found == -1 {
		panic("Item not found")
	}

	return found
}

func (l *List) Clear() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

func (l *List) Concat(k *List) {
	l.Tail.Next, k.Head.Prev = k.Head, l.Tail
	l.Tail = k.Tail
	l.Length += k.Length
}

func (list *List) Map(f func(node *Node)) {
	for node := list.Head; node != nil; node = node.Next {
		n := node.Value.(*Node)
		f(n)
	}
}

func (list *List) Each(f func(node Node) bool) {
	for node := list.Head; node != nil; node = node.Next {
		n := node.Value.(Node)
		f(n)
	}
}
