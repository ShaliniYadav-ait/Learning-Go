package linkedlist

func getDecimalValue(head *ListNode) int {
	
	var s string
	for head != nil{
		s = s + string(head.Val)
		head =  head.Next
	}


	number := anyBaseToDecimal(s, 2)

	return number
}

func anyBaseToDecimal(s string, base int) int {

	val := 0
	valindex := 0
	power := len(s) - 1
digits := []int{'0','1'}

	for i:= 0;i<len(s);i++ {

		for j, conv := range digits {
			if string(conv) == string(s[i]) {
				valindex = j
				break
			}
		}

		val = val + valindex*(recursivePower(base, power))
		power--
	}
	return val
}

func recursivePower(base int, exponent int) int {
	if exponent != 0 {
		return (base * recursivePower(base, exponent-1))
	}
	return 1
}
