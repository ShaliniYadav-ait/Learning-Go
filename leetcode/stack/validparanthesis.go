package stack

//Parenthesis stack
// https://leetcode.com/problems/valid-parentheses/
type Parenthesis struct {
	vals []string
}

// Construct constructor
func Construct() *Parenthesis {
	return &Parenthesis{vals: []string{}}
}

//Push string
func (p *Parenthesis) Push(val string) {
	p.vals = append(p.vals, val)
}

//Pop string
func (p *Parenthesis) Pop() string {
	pop := p.vals[len(p.vals)-1]
	p.vals = p.vals[:len(p.vals)-1]
	return pop
}

//Length function
func (p *Parenthesis) Length() int {
	return len(p.vals)
}

func isValid(s string) bool {

	if len(s) < 2 {
		return false
	}

	paranthesis := Construct()
	for _, val := range s {

		value := string(val)
		if value == "(" || value == "{" || value == "[" {
			paranthesis.Push(value)
			continue
		}

		if paranthesis.Length() == 0{
			return false
		}
		
		pop := paranthesis.Pop()
		switch value {
		case ")":
			if pop != "(" {
				return false
			}
		case "}":
			if pop != "{" {
				return false
			}
		case "]":
			if pop != "[" {
				return false
			}
		default:
			return false
		}
	}

	length := paranthesis.Length()
	return length == 0
}
