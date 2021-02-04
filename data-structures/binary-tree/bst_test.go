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

func TestTraversalAlgorithms_PreOrder(t *testing.T) {
	n := NewNode(1)

	tree := NewTree(n)

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)

	actual := tree.PreOrder(n)
	expected := [...]int{1, 4, 2, 3, 5, 6}

	for i, num := range actual {
		if num != expected[i] {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("PreOrder() = %v, want %v", actual, expected)
			}
		}
	}
}

func TestTraversalAlgorithms_InOrder(t *testing.T) {
	n := NewNode(1)

	tree := NewTree(n)

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)

	actual := tree.InOrder(n)
	expected := [...]int{1, 3, 2, 6, 5, 4}

	for i, num := range actual {
		if num != expected[i] {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("InOrder() = %v, want %v", actual, expected)
			}
		}
	}
}

func TestTraversalAlgorithms_PostOrder(t *testing.T) {
	n := NewNode(1)

	tree := NewTree(n)

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)

	actual := tree.PostOrder(n)
	expected := [...]int{3, 2, 6, 5, 4, 1}

	for i, num := range actual {
		if num != expected[i] {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("PostOrder() = %v, want %v", actual, expected)
			}
		}
	}
}
