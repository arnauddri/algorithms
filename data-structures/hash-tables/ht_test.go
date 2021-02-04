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
	if err != nil || val != "bar" || ht.Size != 5 {
		fmt.Println(val, err)
		t.Error()
	}

	ht.Put("peter", "bob")
	// Test "peter" has been updated
	val, err = ht.Get("peter")
	if err != nil && val == "bob" || ht.Size != 5 {
		fmt.Println(val, err)
		t.Error()
	}

	// Test delete
	ht.Del("peter")
	val, err = ht.Get("peter")
	if val != "" || err == nil || ht.Size != 4 {
		t.Error()
	}

	counter := 0
	f := func(a *item) {
		counter++
	}

	ht.ForEach(f)

	if counter != 4 {
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

func TestHashTable_Resize(t *testing.T) {
	ht := New(1000)
	ht.Put("foo", "bar")
	ht.Put("fiz", "buzz")
	ht.Put("bruce", "wayne")
	ht.Put("peter", "parker")
	ht.Put("clark", "kent")

	actual := ht.Resize(10)

	if actual != nil {
		t.Errorf("resul = %v, want nil", actual)
	}

	actual = ht.Resize(10)
	expected := "current capacity is as same the input number"

	if actual.Error() != expected {
		t.Errorf("resul = %v, want %v", actual, expected)
	}

	actual = ht.Resize(4)
	expected = "there is not enough capacity to hold items. please enter a larger number"

	if actual.Error() != expected {
		t.Errorf("resul = %v, want %v", actual, expected)
	}
}
