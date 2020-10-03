package linkedlist

// https://leetcode.com/problems/linked-list-cycle/
func checkCycleInList(head *ListNode) bool {

	single := head
	double := head

	for double != nil && double.Next != nil {
		single = single.Next
		double = double.Next.Next
		if single == double {
			return true
		}
	}

	return false
}
