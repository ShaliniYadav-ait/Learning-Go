package dailychallenge

// ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	copyList := head
	power := 0

	for copyList != nil {
		power++
		copyList = copyList.Next
	}

	copyList = head
	val := 0
	power--

	for power >= 0{
		val = val + copyList.Val*(recursivePower(power))
		copyList = copyList.Next
		power--
	}

	return val
}

func recursivePower(exponent int) int {
	if exponent != 0 {
		return (2 * recursivePower(exponent-1))
	}
	return 1
}
