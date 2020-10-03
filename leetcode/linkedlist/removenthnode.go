package linkedlist

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummyHead := &ListNode{Next: head}
	endList := head
	beginList := head

	count := 1
	for count < n {
		endList = endList.Next
		count++
	}

	prev := dummyHead
	for endList.Next != nil {
		prev = beginList
		beginList = beginList.Next
		endList = endList.Next
	}

	prev.Next = beginList.Next

	return dummyHead.Next
}
