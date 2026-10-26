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

func TestResize(t *testing.T) {
	ht := New(4)
	keys := []string{"foo", "fiz", "bruce", "peter", "clark", "diana", "barry"}
	for _, k := range keys {
		ht.Put(k, k+"-value")
	}

	for _, newCap := range []int{2, 1000, 7} {
		if err := ht.Resize(newCap); err != nil {
			t.Fatal(err)
		}
		if ht.Capacity != newCap || ht.Size != len(keys) {
			t.Errorf("Resize(%d): Capacity = %d, Size = %d", newCap, ht.Capacity, ht.Size)
		}
		for _, k := range keys {
			if v, err := ht.Get(k); err != nil || v != k+"-value" {
				t.Errorf("Resize(%d): Get(%q) = %v, %v", newCap, k, v, err)
			}
		}
	}

	if ht.Resize(0) == nil {
		t.Error("Resize(0) should fail")
	}
}
