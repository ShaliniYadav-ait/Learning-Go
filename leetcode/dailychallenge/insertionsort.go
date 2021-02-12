package dailychallenge

func insertionSortList(head *ListNode) *ListNode {

	dummyHead := &ListNode{}
	dummyHead.Val = -1 << 63
	dummyHead.Next = head
	last := dummyHead
	curr := head

	for curr != nil {
		if curr.Val >= last.Val {
			last = curr
			curr = curr.Next
			continue
		}

		copyHead := dummyHead.Next
		prev := dummyHead
		
		for copyHead.Val < curr.Val {
			prev = copyHead
			copyHead = copyHead.Next
		}
		last.Next = curr.Next
		prev.Next = curr
		curr.Next = copyHead
		curr = last.Next
	}

	return dummyHead.Next

}
