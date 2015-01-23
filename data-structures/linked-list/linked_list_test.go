package list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Test Prepend/Get
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
	k := NewList()

	k.Append(NewNode(1))
	k.Append(NewNode(2))
	k.Append(NewNode(3))

	if *k.Get(0).Value.(*Node) != *NewNode(1) ||
		*k.Get(1).Value.(*Node) != *NewNode(2) ||
		*k.Get(2).Value.(*Node) != *NewNode(3) {

		fmt.Println(k.Get(0).Value)
		fmt.Println(k.Get(1).Value)
		fmt.Println(k.Get(2).Value)
		t.Error()
	}

	// Test Add
	k.Add(NewNode(8), 1)

	if *k.Get(0).Value.(*Node) != *NewNode(1) ||
		*k.Get(1).Value.(*Node) != *NewNode(8) ||
		*k.Get(2).Value.(*Node) != *NewNode(2) {

		fmt.Println(k.Get(0).Value)
		fmt.Println(k.Get(1).Value)
		fmt.Println(k.Get(2).Value)
		t.Error()
	}

	// Test Concat
	l.Concat(k)
	if l.Len() != 7 {
		t.Error()
	}

	// Test Each
	counter := 0
	f := func(node *Node) {
		//fmt.Println(node.Value)
		counter += node.Value.(int)
	}

	l.Map(f)
	if counter != 20 {
		t.Error()
	}

	// Test Find
	index := l.Find(NewNode(1))
	if index != 3 {
		fmt.Println(index)
		t.Error()
	}

	// Test Clear
	l.Clear()
	if l.Len() != 0 {
		t.Error()
	}

}
