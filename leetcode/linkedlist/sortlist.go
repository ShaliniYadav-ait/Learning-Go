package linkedlist

func sortList(head *ListNode) *ListNode {

	var dummyhead *ListNode
	dummyhead = &ListNode{}
	dummyhead.Next = head
	lessVal := head

	for head != nil {
		prev := head
		head = head.Next
		if head.Val > prev.Val {
			continue
		}
		curr := dummyhead.Next
		for curr.Val <= head.Val {
			lessVal = curr
			curr = curr.Next
		}
		lessVal.Next = head
		head.Next = curr
		head = prev
	}

	return dummyhead.Next
}
