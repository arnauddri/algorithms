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

//every node has a random generated priority
func NewNode(i int) *Node {
	return &Node{Value: i, Priority: rand.Intn(math.MaxUint16)}
}

func (n *Node) Compare(m *Node) int {
	if n.Value < m.Value {
		return -1
	} else if n.Value > m.Value {
		return 1
	} else {
		return 0
	}
}

func (n *Node) ComparePriority(m *Node) int {
	if n.Priority < m.Priority {
		return -1
	} else if n.Priority > m.Priority {
		return 1
	} else {
		return 0
	}
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

func rotateToRight(x *Node, y *Node) {
	x_old := *x
	y_old := *y

	//a ok
	//b
	y.Left = x_old.Right
	if x_old.Right != nil {
		x.Right.Parent = y
	}

	if y_old.Parent != nil {
		if y_old.Parent.Left == y {
			y.Parent.Left = x
		} else {
			y.Parent.Right = x
		}
	}

	x.Parent = y_old.Parent

	//c ok
	//y x
	x.Right = y
	y.Parent = x

}

func rotateToLeft(x *Node, y *Node) {
	x_old := *x
	y_old := *y

	//a ok
	//b
	y.Right = x_old.Left
	if x_old.Left != nil {
		x.Left.Parent = y
	}
	//c ok
	//y x

	if y_old.Parent != nil {
		if y_old.Parent.Left == y {
			y.Parent.Left = x
		} else {
			y.Parent.Right = x
		}
	}

	x.Parent = y_old.Parent

	x.Left = y
	y.Parent = x
}

func (t *Treap) Insert(i int) {
	n := NewNode(i)
	if t.Head == nil {
		t.Head = n
		t.Size++
		return
	}

	h := t.Head
	//walk down with node value as a binary tree
	for {
		if n.Compare(h) == -1 {
			if h.Left == nil {
				h.Left = n
				n.Parent = h
				break
			} else {
				h = h.Left
			}
		} else {
			if h.Right == nil {
				h.Right = n
				n.Parent = h
				break
			} else {
				h = h.Right
			}
		}
	}

	//walk up using random priority as a min heap
	for {
		if n.Parent == nil || n.ComparePriority(n.Parent) == 1 {
			break
		} else {
			if n == n.Parent.Left {
				rotateToRight(n, n.Parent)
			} else {
				rotateToLeft(n, n.Parent)
			}
		}
		if n.Parent == nil {
			t.Head = n
			break
		}
	}

	t.Size++
}

func (t *Treap) Search(i int) *Node {
	h := t.Head
	n := &Node{Value: i}

	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.Right
		case 1:
			h = h.Left
		case 0:
			return h
		default:
			panic("Node not found")
		}
	}
	panic("Node not found")
}
