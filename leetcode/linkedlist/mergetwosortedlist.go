package linkedlist

//  https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var dummyHead *ListNode

	dummyHead = &ListNode{}
	finalList := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			finalList.Next = l1
			l1 = l1.Next
		}else{
			finalList.Next = l2
			l2 = l2.Next
		}
		finalList = finalList.Next
	}

	if l1 != nil{
		finalList.Next = l1
	}

	if l2!= nil{
		finalList.Next = l2
	}

	return dummyHead.Next
}
