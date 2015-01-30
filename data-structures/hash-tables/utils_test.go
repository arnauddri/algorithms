package ht

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	test := "Hello World!"

	if hashCode(test) != 969099747 {
		fmt.Println(hashCode(test))
		t.Error()
	}
}
