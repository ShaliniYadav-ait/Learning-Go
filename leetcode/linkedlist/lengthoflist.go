package linkedlist

func lengthOfList(head *ListNode) (int, bool) {

	lengthMap := map[*ListNode]bool{}
	count := 0
	hasCycle := false

	for head != nil {
		if lengthMap[head] {
			hasCycle = true
			break
		}
		count++
		lengthMap[head] = true
		head = head.Next
	}
	return count, hasCycle
}
