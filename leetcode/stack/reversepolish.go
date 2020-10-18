package stack

//https://leetcode.com/problems/evaluate-reverse-polish-notation/
import (
	"strconv"
)

//RevStack structure
type RevStack struct {
	vals []int
}

// constructor
func constructor() *RevStack {
	return &RevStack{vals: []int{}}
}

// PushRev function
func (rev *RevStack) PushRev(val int) {
	rev.vals = append(rev.vals, val)
}

// PopRev function
func (rev *RevStack) PopRev() int {
	last := rev.vals[len(rev.vals)-1]

	rev.vals = rev.vals[:len(rev.vals)-1]
	return last
}

func evalRPN(tokens []string) int {

	stack := constructor()
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "/" || tokens[i] == "*" {
			num2 := stack.PopRev()
		    num1 := stack.PopRev()
			var val int
			switch tokens[i] {
			case "+":
				val = num1 + num2
			case "-":
				val = num1 - num2
			case "/":
				val = num1 / num2
			case "*":
				val = num1 * num2
			}
			stack.PushRev(val)
		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack.PushRev(num)
		}
	}
return stack.PopRev()
}
