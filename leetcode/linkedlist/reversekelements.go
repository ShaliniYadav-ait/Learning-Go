package linkedlist

// https://leetcode.com/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *ListNode, k int) *ListNode {

	
	dummyHead := &ListNode{
		Next : head,
	}
	copyList := head

	if k < 2 {
		return head
	}

	count := 1
	for copyList.Next != nil {
		copyList = copyList.Next
		count++
	}

	groups := count / k
	prev := dummyHead

	for groups != 0 {
		reversek(prev, k)

		for i := 1; i <= k; i++ {
			prev = prev.Next
		}
		groups--
	}
	return dummyHead.Next
}

func reversek(prev *ListNode, k int) {
	
	headList := prev.Next
	n := 1
	for n < k {
		temp := headList.Next
		headList.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
		n++
	}
}
