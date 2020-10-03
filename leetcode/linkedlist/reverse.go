package linkedlist

// ListNode : https://leetcode.com/problems/reverse-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil{
		return head
	}

	var dummyHead *ListNode
	dummyHead = &ListNode{}
	dummyHead.Next = head

	reverse:= dummyHead.Next
	for reverse.Next != nil{
		temp := reverse.Next
		reverse.Next = temp.Next
		temp.Next = dummyHead.Next
		dummyHead.Next = temp
	} 
	return dummyHead.Next
}
