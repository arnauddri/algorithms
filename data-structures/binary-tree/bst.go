package bst

import ()

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(i int) *Node {
	return &Node{Value: i}
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

type Tree struct {
	Head *Node
	Size int
}

func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{Head: n, Size: 1}
}

func (t *Tree) Insert(i int) {
	n := &Node{Value: i}
	if t.Head == nil {
		t.Head = n
		t.Size++
		return
	}

	h := t.Head

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
	t.Size++
}

func (t *Tree) Search(i int) *Node {
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

// returns true if a node with value i was found
// and deleted and returns false otherwise
func (t *Tree) Delete(i int) bool {
	var parent *Node

	h := t.Head
	n := &Node{Value: i}
	for h != nil {
		switch n.Compare(h) {
		case -1:
			parent = h
			h = h.Left
		case 1:
			parent = h
			h = h.Right
		case 0:
			if h.Left != nil {
				right := h.Right
				h.Value = h.Left.Value
				h.Left = h.Left.Left
				h.Right = h.Left.Right

				if right != nil {
					subTree := &Tree{Head: h}
					IterOnTree(right, func(n *Node) {
						subTree.Insert(n.Value)
					})
				}
				t.Size--
				return true
			}

			if h.Right != nil {
				h.Value = h.Right.Value
				h.Left = h.Right.Left
				h.Right = h.Right.Right

				t.Size--
				return true
			}

			if parent == nil {
				t.Head = nil
				t.Size--
				return true
			}

			if parent.Left == n {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			t.Size--
			return true
		}
	}
	return false
}

func IterOnTree(n *Node, f func(*Node)) bool {
	if n == nil {
		return true
	}
	if !IterOnTree(n.Left, f) {
		return false
	}

	f(n)

	return IterOnTree(n.Right, f)
}
