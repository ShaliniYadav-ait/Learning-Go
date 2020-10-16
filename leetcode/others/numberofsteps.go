package others

//https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
			count++
		} else {
			num = num - 1
			count++

			if num != 0 {
				num = num / 2
				count++
			}
		}
	}
	return count
}
