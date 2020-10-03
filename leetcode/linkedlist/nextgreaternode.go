package linkedlist

type listValIndex struct {
	Val   int
	Index int
	Next  *listValIndex
}

func nextLargerNodes(head *ListNode) []int {

	var stack *listValIndex
	index := 0
	result := []int{}

	for head != nil {

		result = append(result, 0)

		for stack != nil && head.Val > stack.Val {
			result[stack.Index] = head.Val
			stack = stack.Next
		}

		slicenode := &listValIndex{
			Val:   head.Val,
			Index: index,
			Next:  stack,
		}

		stack = slicenode
		head = head.Next
		index++
	}

	return result
}
