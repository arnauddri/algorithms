// Treap: a binary search tree on Value that is also a min-heap on a random
// Priority, which keeps the tree balanced in expectation.
// https://en.wikipedia.org/wiki/Treap
package treap

import (
	"math"
	"math/rand"
)

type Node struct {
	Value    int
	Priority int
	Parent   *Node
	Left     *Node
	Right    *Node
}

// NewNode returns a node with a random priority.
func NewNode(i int) *Node {
	return &Node{Value: i, Priority: rand.Intn(math.MaxUint16)}
}

func (n *Node) Compare(m *Node) int {
	if n.Value < m.Value {
		return -1
	} else if n.Value > m.Value {
		return 1
	}
	return 0
}

func (n *Node) ComparePriority(m *Node) int {
	if n.Priority < m.Priority {
		return -1
	} else if n.Priority > m.Priority {
		return 1
	}
	return 0
}

type Treap struct {
	Head *Node
	Size int
}

func NewTreap(n *Node) *Treap {
	if n == nil {
		return &Treap{}
	}
	return &Treap{Head: n, Size: 1}
}

// rotateUp moves x above its parent y, preserving in-order.
func (t *Treap) rotateUp(x *Node) {
	y := x.Parent
	if y == nil {
		return
	}

	if x == y.Left {
		y.Left = x.Right
		if x.Right != nil {
			x.Right.Parent = y
		}
		x.Right = y
	} else {
		y.Right = x.Left
		if x.Left != nil {
			x.Left.Parent = y
		}
		x.Left = y
	}

	x.Parent = y.Parent
	y.Parent = x

	if x.Parent == nil {
		t.Head = x
	} else if x.Parent.Left == y {
		x.Parent.Left = x
	} else {
		x.Parent.Right = x
	}
}

// Insert adds a node as a leaf (BST rule), then rotates it up until its
// priority is not lower than its parent's (heap rule). Duplicates go right.
func (t *Treap) Insert(i int) {
	n := NewNode(i)
	t.Size++

	if t.Head == nil {
		t.Head = n
		return
	}

	h := t.Head
	for {
		if n.Compare(h) == -1 {
			if h.Left == nil {
				h.Left = n
				n.Parent = h
				break
			}
			h = h.Left
		} else {
			if h.Right == nil {
				h.Right = n
				n.Parent = h
				break
			}
			h = h.Right
		}
	}

	for n.Parent != nil && n.ComparePriority(n.Parent) == -1 {
		t.rotateUp(n)
	}
}

// Search returns the node holding i, or nil if absent.
func (t *Treap) Search(i int) *Node {
	h := t.Head
	n := &Node{Value: i}

	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.Right
		case 1:
			h = h.Left
		default:
			return h
		}
	}
	return nil
}

// Delete rotates the node holding i down to a leaf, always lifting the child
// with the lower priority, then detaches it. Returns false if i is absent.
func (t *Treap) Delete(i int) bool {
	n := t.Search(i)
	if n == nil {
		return false
	}

	for n.Left != nil || n.Right != nil {
		if n.Right == nil || (n.Left != nil && n.Left.ComparePriority(n.Right) == -1) {
			t.rotateUp(n.Left)
		} else {
			t.rotateUp(n.Right)
		}
	}

	switch {
	case n.Parent == nil:
		t.Head = nil
	case n.Parent.Left == n:
		n.Parent.Left = nil
	default:
		n.Parent.Right = nil
	}
	n.Parent = nil
	t.Size--

	return true
}
