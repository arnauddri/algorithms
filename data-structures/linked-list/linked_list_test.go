package list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Test Prepend
	l := NewList()

	l.Prepend(NewNode(1))
	l.Prepend(NewNode(2))
	l.Prepend(NewNode(3))

	if *l.Get(0).Value.(*Node) != *NewNode(3) ||
		*l.Get(1).Value.(*Node) != *NewNode(2) ||
		*l.Get(2).Value.(*Node) != *NewNode(1) {

		fmt.Println(l.Get(0).Value)
		fmt.Println(l.Get(1).Value)
		fmt.Println(l.Get(2).Value)
		t.Error()
	}

	// Test Append
	l = NewList()

	l.Append(NewNode(1))
	l.Append(NewNode(2))
	l.Append(NewNode(3))

	if *l.Get(0).Value.(*Node) != *NewNode(1) ||
		*l.Get(1).Value.(*Node) != *NewNode(2) ||
		*l.Get(2).Value.(*Node) != *NewNode(3) {

		fmt.Println(l.Get(0).Value)
		fmt.Println(l.Get(1).Value)
		fmt.Println(l.Get(2).Value)
		t.Error()
	}

	l.Add(NewNode(8), 1)

	if *l.Get(0).Value.(*Node) != *NewNode(1) ||
		*l.Get(1).Value.(*Node) != *NewNode(8) ||
		*l.Get(2).Value.(*Node) != *NewNode(2) {

		fmt.Println(l.Get(0).Value)
		fmt.Println(l.Get(1).Value)
		fmt.Println(l.Get(2).Value)
		t.Error()
	}
}
