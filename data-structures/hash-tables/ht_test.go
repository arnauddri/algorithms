package ht

import (
	"fmt"
	"testing"
)

func TestHt(t *testing.T) {
	ht := New(1000)
	ht.Put("foo", "bar")
	ht.Put("fiz", "buzz")
	ht.Put("bruce", "wayne")
	ht.Put("peter", "parker")
	ht.Put("clark", "kent")

	// Test simple get
	val, err := ht.Get("foo")
	if err != nil && val == "bar" {
		fmt.Println(val, err)
		t.Error()
	}

	ht.Put("peter", "bob")
	// Test "peter" has been updated
	val, err = ht.Get("peter")
	if err != nil && val == "bob" {
		fmt.Println(val, err)
		t.Error()
	}

}

func TestHash(t *testing.T) {
	test := "Hello World!"

	if hashCode(test) != 969099747 {
		fmt.Println(hashCode(test))
		t.Error()
	}
}
