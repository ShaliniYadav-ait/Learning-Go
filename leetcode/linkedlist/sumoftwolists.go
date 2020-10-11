package linkedlist

func sumOfTwoLinkedLists(head1 *ListNode, head2 *ListNode) *ListNode {

	dummyhead := &ListNode{}
	head := dummyhead
	num1, num2, sum, count, val := 0, 0, 0, 0, 1

	num1 = findNumber(head1)
	num2 = findNumber(head2)

	sum = num1 + num2
	count = 1

	for {
		val = sum % 10
		sum = sum / 10
		if val != 0 || count != 0 {
			node := &ListNode{Val: val}
			head.Next = node
			head = head.Next
			count = 0
			continue
		} 
			break
	}

	return dummyhead.Next
}

func findNumber(head *ListNode) int {

	num, count := 0, 0

	for head != nil {
		num = num + findPower(head.Val, count)
		head = head.Next
		count++
	}
	return num
}

func findPower(num int, power int) int {

	mul := 1
	for power != 0 {
		mul = mul * 10
		power--
	}

	return num * mul
}
