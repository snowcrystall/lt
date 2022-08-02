package exercise

import (
	"errors"
)

//实现一个栈，能够O(1)时间获取栈中元素的最大值

type MaxStach interface {
	Max() (int, error)
	Pop() (int, error)
	Push(int)
}

type stack struct {
	a []int
	b []int
}

func NewStack() MaxStach {
	return &stack{[]int{}, []int{}}
}
func (s *stack) Max() (int, error) {
	if len(s.b) == 0 {
		return -1, errors.New("empty")
	}
	return s.b[len(s.b)-1], nil
}

func (s *stack) Pop() (int, error) {
	if len(s.a) == 0 {
		return -1, errors.New("empty")
	}
	r := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]

	if r == s.b[len(s.b)-1] {
		s.b = s.b[:len(s.b)-1]
	}
	return r, nil
}

func (s *stack) Push(x int) {
	s.a = append(s.a, x)
	if len(s.b) == 0 {
		s.b = append(s.b, x)
	}
	if x >= s.b[len(s.b)-1] {
		s.b = append(s.b, x)
	}
}
