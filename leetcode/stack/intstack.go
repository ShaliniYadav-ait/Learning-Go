package stack

import (
	"errors"
)

// Stack type for Stack implementation
type Stack struct {
	vals []int
}

func initialize() *Stack {
	return &Stack{vals: []int{}}
}

// Push function for Stack
func (s *Stack) Push(val int) {
	s.vals = append(s.vals, val)
}

// Pop function for Stack
func (s *Stack) Pop() (int, error) {
	if s.Length() == 0 {
		return 0, errors.New("Stack is empty! Cannot Pop")
	}

	result := s.vals[s.Length()-1]
	s.vals = s.vals[:(s.Length() - 2)]
	return result, nil
}

// Length function for Stack
func (s *Stack) Length() int {
	return len(s.vals)
}
