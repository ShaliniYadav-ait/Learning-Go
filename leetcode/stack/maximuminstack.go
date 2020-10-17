package stack

import (
	"errors"
)

// MaxStack Implementation
type MaxStack struct {
		vals []int
		max  []int
}

func create() *MaxStack {
	return &MaxStack{vals: []int{}, max: []int{}}
}

//PushMax MaxStack
func (m *MaxStack) PushMax(val int) {
	if len(m.max) != 0 {
		max := m.max[len(m.max)-1]
		if max > val {
			m.max = append(m.max, max)
		} else {
			m.max = append(m.max, val)
		}
	} else {
		m.max = append(m.max, val)
	}

	m.vals = append(m.vals, val)
}

//PopMax MaxStack
func (m *MaxStack) PopMax() (int, error) {

	val := 0
	if len(m.vals) != 0 {

		val = m.vals[len(m.vals)-1]
		m.vals = m.vals[:len(m.vals)-2]
		m.max = m.max[:len(m.max)-2]

	} else {
		return -1, errors.New("Stack is empty! cannot Pop")
	}
	return val, nil
}

//GetMaximum MaxStack
func (m *MaxStack) GetMaximum() (int, error) {

	maximum := 0
	if len(m.max) != 0 {
		maximum = m.max[len(m.max)-1]
	} else {
		return -1, errors.New("There are no values in the stack")
	}
	return maximum,nil
}
