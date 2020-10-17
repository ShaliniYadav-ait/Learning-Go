package stack

//MinStack type declaration 
//https://leetcode.com/problems/min-stack/
type MinStack struct {
	vals []int
	min  []int
}

//Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{vals: []int{}, min: []int{}}
}

//Push stack
func (m *MinStack) Push(val int) {
	if len(m.min) != 0 {
		min := m.min[len(m.min)-1]
		if min < val {
			m.min = append(m.min, min)
		} else {
			m.min = append(m.min, val)
		}
	} else {
		m.min = append(m.min, val)
	}

	m.vals = append(m.vals, val)
}

//Pop stack
func (m *MinStack) Pop() {

	if len(m.vals) != 0 {
		m.vals = m.vals[:len(m.vals)-1]
		m.min = m.min[:len(m.min)-1]
	}
}

//Top stack
func (m *MinStack) Top() int {
	top := 0
	if len(m.vals) != 0 {
		top = m.vals[len(m.vals)-1]
	} else {
		return -1
	}
	return top
}

//GetMin stack
func (m *MinStack) GetMin() int {
	minimum := 0
	if len(m.min) != 0 {
		minimum = m.min[len(m.min)-1]
	} else {
		return -1
	}
	return minimum
}
