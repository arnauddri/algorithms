package stack

type Stack struct {
	stack []interface{}
	len   int
}

func New() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0

	return s
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) isEmpty() bool {
	return s.len == 0
}

func (s *Stack) Pop() (el interface{}) {
	el, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

func (s *Stack) Push(el interface{}) {
	prepend := make([]interface{}, 1)
	prepend[0] = el
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Peek() interface{} {
	return s.stack[0]
}
