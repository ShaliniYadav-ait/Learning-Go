package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	
	if head == nil {
		return head
	}
	dummyheadodd := &ListNode{}
	dummyheadeven := &ListNode{}

	curr := head
	count := 0

	for curr != nil {
		count++
		curr = curr.Next
	}

	even := dummyheadeven
	odd := dummyheadodd
	curr = head
	for i := 1; i <= count; i++ {
		if i%2 == 0 {
			even.Next = curr
			even = even.Next
			curr = curr.Next
		}

		if i%2 != 0 {
			odd.Next = curr
			odd = odd.Next
			curr = curr.Next
		}
	}

	odd.Next = dummyheadeven.Next
	odd = dummyheadodd.Next
	for count != 1 {
		odd = odd.Next
		count--
	}
	odd.Next = nil

	return dummyheadodd.Next
}
