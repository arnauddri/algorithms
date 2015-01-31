package ht

import (
	"fmt"
	"testing"
)

func TestHt(t *testing.T) {
	ht := New(1000)
	ht.Put("foo", "bar")
	//ht.Put("foo", "bar")
	val, _ := ht.Get("foo")

	fmt.Println(val)
}

func TestHash(t *testing.T) {
	test := "Hello World!"

	if hashCode(test) != 969099747 {
		fmt.Println(hashCode(test))
		t.Error()
	}
}
