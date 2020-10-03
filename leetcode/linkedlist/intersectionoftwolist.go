package linkedlist

// https://leetcode.com/problems/intersection-of-two-linked-lists/

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lengthOfHeadA := findLength(headA)
	lengthOfHeadB := findLength(headB)
	headOfListA := headA
	headOfListB := headB
	diff := 0

	if lengthOfHeadA > lengthOfHeadB {
		diff = lengthOfHeadA - lengthOfHeadB
		for diff != 0 {
			headOfListA = headOfListA.Next
			diff--
		}
	} else {
		diff = lengthOfHeadB - lengthOfHeadA
		for diff != 0 {
			headOfListB = headOfListB.Next
			diff--
		}
	}

	for headOfListA != nil || headOfListB != nil {
		if headOfListA == headOfListB {
			return headOfListA
		}
		headOfListA = headOfListA.Next
		headOfListB = headOfListB.Next
	}
	return nil

}

func findLength(Head *ListNode) int {

	if Head == nil {
		return 0
	}

	n := 1
	copy := Head

	for copy.Next != nil {
		copy = copy.Next
		n++
	}

	return n
}
