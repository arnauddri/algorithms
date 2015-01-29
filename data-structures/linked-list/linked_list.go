package list

import (
	"errors"
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

func (l *List) Add(value interface{}, index int) error {
	if index > l.Len() {
		return errors.New("Index out of range")
	}

	node := NewNode(value)

	if l.Len() == 0 || index == 0 {
		l.Prepend(value)
		return nil
	}

	if l.Len()-1 == index {
		l.Append(value)
		return nil
	}

	nextNode, _ := l.Get(index)
	prevNode := nextNode.Prev

	prevNode.Next = node
	node.Prev = prevNode

	nextNode.Prev = node
	node.Next = nextNode

	l.Length++

	return nil
}

func (l *List) Remove(value interface{}) error {
	if l.Len() == 0 {
		return errors.New("Empty list")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	found := 0
	for n := l.Head; n != nil; n = n.Next {

		if *n.Value.(*Node) == value && found == 0 {
			n.Next.Prev, n.Prev.Next = n.Prev, n.Next
			l.Length--
			found++
		}
	}

	if found == 0 {
		return errors.New("Node not found")
	}

	return nil
}

func (l *List) Get(index int) (*Node, error) {
	if index > l.Len() {
		return nil, errors.New("Index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

func (l *List) Find(node *Node) (int, error) {
	if l.Len() == 0 {
		return 0, errors.New("Empty list")
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
		return 0, errors.New("Item not found")
	}

	return found, nil
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

func (list *List) Each(f func(node Node)) {
	for node := list.Head; node != nil; node = node.Next {
		f(*node)
	}
}
