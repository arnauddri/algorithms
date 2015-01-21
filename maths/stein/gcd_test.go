package gcd

import (
	"testing"
)

func TestRecurse(t *testing.T) {
	if recurse(14, 7) != 7 ||
		recurse(4, 2) != 2 ||
		recurse(31, 2) != 1 ||
		recurse(33, 11) != 11 {
		t.Error()
	}
}
