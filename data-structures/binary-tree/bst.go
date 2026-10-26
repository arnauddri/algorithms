package bst

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
	h := t.Head
	n := &Node{Value: i}
	for h != nil {
		switch n.Compare(h) {
		case -1:
			h = h.Left
		case 1:
			h = h.Right
		case 0:
			// Two children: copy the in-order successor up, then remove it
			if h.Left != nil && h.Right != nil {
				succ := h.Right
				for succ.Left != nil {
					succ = succ.Left
				}
				h.Value = succ.Value
				h = succ
			}

			child := h.Left
			if child == nil {
				child = h.Right
			}
			t.replace(h, child)
			t.Size--
			return true
		}
	}
	return false
}

// replace swaps n for child in n's parent (or at the root)
func (t *Tree) replace(n, child *Node) {
	if child != nil {
		child.Parent = n.Parent
	}
	switch {
	case n.Parent == nil:
		t.Head = child
	case n.Parent.Left == n:
		n.Parent.Left = child
	default:
		n.Parent.Right = child
	}
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

// PreOrder returns the values of the subtree rooted at n in pre-order
func (t *Tree) PreOrder(n *Node) []int {
	if n == nil {
		return nil
	}
	nodes := []int{n.Value}
	nodes = append(nodes, t.PreOrder(n.Left)...)
	return append(nodes, t.PreOrder(n.Right)...)
}

// InOrder returns the values of the subtree rooted at n in in-order
func (t *Tree) InOrder(n *Node) []int {
	if n == nil {
		return nil
	}
	nodes := t.InOrder(n.Left)
	nodes = append(nodes, n.Value)
	return append(nodes, t.InOrder(n.Right)...)
}

// PostOrder returns the values of the subtree rooted at n in post-order
func (t *Tree) PostOrder(n *Node) []int {
	if n == nil {
		return nil
	}
	nodes := t.PostOrder(n.Left)
	nodes = append(nodes, t.PostOrder(n.Right)...)
	return append(nodes, n.Value)
}
