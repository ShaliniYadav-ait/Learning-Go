package queue

import (
	"errors"
)

//Queue of integer
type Queue struct {
	vals []int
}

func initialize() *Queue {
	return &Queue{vals: []int{}}
}

// Push integer
func (q *Queue) Push(val int) {
	q.vals = append(q.vals, val)
}

//Pop integer
func (q *Queue) Pop() (int, error) {
	if q.Length() == 0{
		return 0, errors.New("Empty Queue")
	}

	val := q.vals[0]
	q.vals = q.vals[1:]
	return val,nil
}

//Length of queue
func (q *Queue) Length() int {
	return len(q.vals)
}
