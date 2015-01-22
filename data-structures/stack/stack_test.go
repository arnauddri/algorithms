package stack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()

	if !s.isEmpty() ||
		s.len != 0 ||
		s.Len() != 0 {
		t.Error()
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.stack[0] != 3 ||
		s.stack[1] != 2 ||
		s.stack[2] != 1 {
		fmt.Println(s.stack)
		t.Error()
	}

	if s.Len() != 3 {
		t.Error()
	}

	a := s.Pop()

	if a != 3 || s.Len() != 2 {
		t.Error()
	}

	b := s.Peek()

	if b != 2 {
		t.Error()
	}
}
