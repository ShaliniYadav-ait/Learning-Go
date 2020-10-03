package linkedlist

// https://leetcode.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {

	if head != nil {
		return nil
	}
	
	cycle := false
	single := head
	double := head.Next

	for double != nil && double.Next != nil {
		single = single.Next
		double = double.Next.Next
		if single == double {
			cycle = true
			break
		}

	}

	if cycle == false {
		return nil
	}

	perimeter := 1
	double = double.Next
	for single != double {
		double = double.Next
		perimeter++
	}

	single = head
	double = head
	count := 1
	for count <= perimeter {
		double = double.Next
		count++
	}

	for single != double {
		single = single.Next
		double = double.Next
	}

	return single
}
