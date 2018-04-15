package treap

import (
	"testing"
)

func TestTreap(t *testing.T) {

	n := NewNode(1)
	m := NewNode(2)

	// Test compare
	if n.Compare(m) != -1 || m.Compare(n) != 1 || n.Compare(n) != 0 {
		t.Error()
	}

	Treap := NewTreap(n)

	Treap.Insert(4)
	Treap.Insert(2)
	Treap.Insert(5)
	Treap.Insert(3)
	Treap.Insert(6)

	if Treap.Size != 6 {
		t.Error()
	}

	five := Treap.Search(5)

	if five.Value != 5 {
		t.Error()
	}

	// Treap.Delete(5)
	//
	// if Treap.Size != 5 {
	// 	t.Error()
	// }
	//
	// four := *Treap.Search(4)
	// if four.Right.Value != 6 ||
	// 	four.Left.Value != 2 ||
	// 	four.Parent.Value != 1 {
	// 	fmt.Println(*Treap.Search(4))
	// 	t.Error()
	// }
}
