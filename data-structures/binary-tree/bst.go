package bst

import ()

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

func (n *Node) Compare(than *Node) int {
	if n.Value < than.Value {
		return -1
	} else if n.Value > than.Value {
		return 1
	} else {
		return 0
	}
}

type Tree struct {
	Head *Node
	Size int
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
			if n.Left == nil {
				n.Left = h
				break
			} else {
				n = n.Left
			}
		} else {
			if n.Right == nil {
				n.Right = h
				break
			} else {
				n = n.Right
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
