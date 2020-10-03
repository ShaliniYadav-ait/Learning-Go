package linkedlist

//package linkedlist : https://leetcode.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	var dummyHead *ListNode
	dummyHead = &ListNode{}
	dummyHead.Next = head
	beforeHead := dummyHead

	if m == n {
		return head
	}

	count := 1
	for count != m {
		beforeHead = beforeHead.Next
		count++
	}

	reverse := beforeHead.Next
	for m < n {
		temp := reverse.Next
		reverse.Next = temp.Next
		temp.Next = beforeHead.Next
		beforeHead.Next = temp
		m++
	}
	
	return dummyHead.Next
}
