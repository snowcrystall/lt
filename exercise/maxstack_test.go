package exercise

import "testing"

func TestMaxstack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(2)
	s.Push(2)
	s.Push(3)
	s.Push(2)
	s.Push(2)
	s.Push(4)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	if max, _ := s.Max(); max != 2 {
		t.Error(s.Max())
	}
}
