package treap

import (
	"math/rand"
	"sort"
	"testing"
)

// inOrder returns the values in BST order and fails if the heap or parent
// invariants are broken anywhere in the tree.
func inOrder(t *testing.T, n *Node, parent *Node) []int {
	if n == nil {
		return nil
	}
	if n.Parent != parent {
		t.Fatalf("node %d has wrong parent", n.Value)
	}
	if parent != nil && n.ComparePriority(parent) == -1 {
		t.Fatalf("node %d has lower priority than its parent", n.Value)
	}
	out := inOrder(t, n.Left, n)
	out = append(out, n.Value)
	return append(out, inOrder(t, n.Right, n)...)
}

func checkInvariants(t *testing.T, tr *Treap) {
	values := inOrder(t, tr.Head, nil)
	if len(values) != tr.Size {
		t.Fatalf("size %d, but %d nodes", tr.Size, len(values))
	}
	if !sort.IntsAreSorted(values) {
		t.Fatalf("not in BST order: %v", values)
	}
}

func TestTreap(t *testing.T) {
	n := NewNode(1)
	m := NewNode(2)

	if n.Compare(m) != -1 || m.Compare(n) != 1 || n.Compare(n) != 0 {
		t.Error()
	}

	tr := NewTreap(n)

	tr.Insert(4)
	tr.Insert(2)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(6)

	if tr.Size != 6 {
		t.Error()
	}
	checkInvariants(t, tr)

	if five := tr.Search(5); five == nil || five.Value != 5 {
		t.Error()
	}
	if tr.Search(42) != nil {
		t.Error("Search must return nil for an absent value")
	}
}

func TestDelete(t *testing.T) {
	tr := NewTreap(nil)
	for _, v := range []int{1, 4, 2, 5, 3, 6} {
		tr.Insert(v)
	}

	if !tr.Delete(5) || tr.Size != 5 || tr.Search(5) != nil {
		t.Error()
	}
	checkInvariants(t, tr)

	if tr.Delete(5) {
		t.Error("deleting an absent value must return false")
	}

	// Deleting the root must keep the treap valid.
	if !tr.Delete(tr.Head.Value) {
		t.Error()
	}
	checkInvariants(t, tr)

	for _, v := range inOrder(t, tr.Head, nil) {
		tr.Delete(v)
	}
	if tr.Size != 0 || tr.Head != nil {
		t.Error()
	}
}

func TestRandomOperations(t *testing.T) {
	tr := NewTreap(nil)
	values := rand.Perm(500)

	for _, v := range values {
		tr.Insert(v)
	}
	checkInvariants(t, tr)

	for _, v := range values[:250] {
		if !tr.Delete(v) {
			t.Errorf("failed to delete %d", v)
		}
	}
	checkInvariants(t, tr)

	for _, v := range values[:250] {
		if tr.Search(v) != nil {
			t.Errorf("%d still present", v)
		}
	}
	for _, v := range values[250:] {
		if tr.Search(v) == nil {
			t.Errorf("%d missing", v)
		}
	}
}
