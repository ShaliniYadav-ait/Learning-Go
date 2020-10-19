package stack

import (
	"strings"
)

// Simplify structure
//https://leetcode.com/problems/simplify-path/
type Simplify struct {
	vals []string
}

// constructSimplify stack
func constructSimplify() *Simplify {
	return &Simplify{vals: []string{}}
}

//Push Path
func (s *Simplify) Push(val string) {
	s.vals = append(s.vals, val)
}

//Pop Path
func (s *Simplify) Pop() string {
	element := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return element
}

//Delete Path
func (s *Simplify) Delete() {
	if s.Length() != 0 {
		s.vals = s.vals[:len(s.vals)-1]
	}
}

// Length of stack
func (s *Simplify) Length() int {
	return len(s.vals)
}

func simplifyPath(path string) string {

	if path == "" {
		return ""
	}

	s := strings.Split(path, "/")
	simplify := constructSimplify()
	result := ""
	tot := 0

	for i := 0; i < len(s); i++ {
		if s[i] == "" || s[i] == "." {
			continue
		}

		if s[i] == ".." {
			simplify.Delete()
			if tot != 0 {
				tot--
			}
			continue
		}

		simplify.Push(s[i])
		tot++

	}

	if simplify.Length() == 0 {
		return "/"
	}

	for tot != 0 {
		element := simplify.Pop()
		result = "/" + element + result
		tot--
	}

	return result
}
