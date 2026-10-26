package bst

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	n := NewNode(1)
	m := NewNode(2)

	// Test compare
	if n.Compare(m) != -1 || m.Compare(n) != 1 || n.Compare(n) != 0 {
		fmt.Println(n.Compare(m))
		t.Error()
	}

	tree := NewTree(n)

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)

	if tree.Size != 6 {
		fmt.Println(tree.Size)
		t.Error()
	}

	five := tree.Search(5)

	if five.Value != 5 ||
		five.Parent.Value != 4 ||
		five.Right.Value != 6 ||
		five.Left != nil {
		fmt.Println(*tree.Search(5))
		t.Error()
	}

	tree.Delete(5)

	if tree.Size != 5 {
		t.Error()
	}

	four := *tree.Search(4)
	if four.Right.Value != 6 ||
		four.Left.Value != 2 ||
		four.Parent.Value != 1 {
		fmt.Println(*tree.Search(4))
		t.Error()
	}
}

func TestDelete(t *testing.T) {
	values := []int{50, 30, 70, 20, 40, 60, 80, 35, 45, 65}
	for _, del := range values {
		tree := NewTree(nil)
		for _, v := range values {
			tree.Insert(v)
		}
		if !tree.Delete(del) {
			t.Errorf("Delete(%d) = false", del)
		}

		var expected []int
		for _, v := range []int{20, 30, 35, 40, 45, 50, 60, 65, 70, 80} {
			if v != del {
				expected = append(expected, v)
			}
		}
		if got := tree.InOrder(tree.Head); !reflect.DeepEqual(got, expected) {
			t.Errorf("after Delete(%d): InOrder() = %v, want %v", del, got, expected)
		}
		if tree.Size != len(expected) {
			t.Errorf("after Delete(%d): Size = %d, want %d", del, tree.Size, len(expected))
		}
		IterOnTree(tree.Head, func(n *Node) {
			if n.Left != nil && n.Left.Parent != n || n.Right != nil && n.Right.Parent != n {
				t.Errorf("after Delete(%d): broken parent pointer at %d", del, n.Value)
			}
		})
	}

	tree := NewTree(NewNode(1))
	if tree.Delete(2) {
		t.Error("Delete of missing value returned true")
	}
	tree.Delete(1)
	if tree.Head != nil || tree.Size != 0 {
		t.Error("deleting the only node should empty the tree")
	}
}

func TestTraversals(t *testing.T) {
	tree := NewTree(NewNode(1))
	for _, v := range []int{4, 2, 5, 3, 6} {
		tree.Insert(v)
	}

	cases := []struct {
		name     string
		got      []int
		expected []int
	}{
		{"PreOrder", tree.PreOrder(tree.Head), []int{1, 4, 2, 3, 5, 6}},
		{"InOrder", tree.InOrder(tree.Head), []int{1, 2, 3, 4, 5, 6}},
		{"PostOrder", tree.PostOrder(tree.Head), []int{3, 2, 6, 5, 4, 1}},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.got, c.expected) {
			t.Errorf("%s() = %v, want %v", c.name, c.got, c.expected)
		}
	}
}
