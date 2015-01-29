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

	zero := *slice(l.Get(0))[0].(*Node).Value.(*Node)
	one := *slice(l.Get(1))[0].(*Node).Value.(*Node)
	two := *slice(l.Get(2))[0].(*Node).Value.(*Node)

	if zero != *NewNode(3) ||
		one != *NewNode(2) ||
		two != *NewNode(1) {

		fmt.Println(*one.Value.(*Node), *NewNode(2))
		fmt.Println(zero.Value)
		fmt.Println(one.Value)
		fmt.Println(two.Value)
		t.Error()
	}

	// Test Append
	k := NewList()

	k.Append(NewNode(1))
	k.Append(NewNode(2))
	k.Append(NewNode(3))

	zero = *slice(k.Get(0))[0].(*Node).Value.(*Node)
	one = *slice(k.Get(1))[0].(*Node).Value.(*Node)
	two = *slice(k.Get(2))[0].(*Node).Value.(*Node)

	if zero != *NewNode(1) ||
		one != *NewNode(2) ||
		two != *NewNode(3) {

		fmt.Println(zero.Value)
		fmt.Println(one.Value)
		fmt.Println(two.Value)
		t.Error()
	}

	// Test Add
	k.Add(NewNode(8), 1)

	zero = *slice(k.Get(0))[0].(*Node).Value.(*Node)
	one = *slice(k.Get(1))[0].(*Node).Value.(*Node)
	two = *slice(k.Get(2))[0].(*Node).Value.(*Node)

	if zero != *NewNode(1) ||
		one != *NewNode(8) ||
		two != *NewNode(2) {

		fmt.Println(zero.Value)
		fmt.Println(one.Value)
		fmt.Println(two.Value)

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
		counter += node.Value.(int)
	}

	l.Map(f)
	if counter != 20 {
		t.Error()
	}

	// Test Find
	index, _ := l.Find(NewNode(1))
	if index != 3 {
		fmt.Println(index)
		t.Error()
	}

	// Test Remove
	l.Remove(*NewNode(8))

	counter = 0
	l.Map(func(n *Node) {
		counter += n.Value.(int)
	})

	if counter != 12 {
		t.Error()
	}

	// Test Clear
	l.Clear()
	if l.Len() != 0 {
		t.Error()
	}

}

func slice(args ...interface{}) []interface{} {
	return args
}
