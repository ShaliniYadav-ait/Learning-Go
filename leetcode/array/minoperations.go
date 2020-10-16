package array

//https://leetcode.com/problems/minimum-operations-to-make-array-equal/
func minOperations(n int) int {

	avg, result := 0, 0

	for i := 0; i < n; i++ {
		avg = avg + ((2 * i) + 1)
	}

	avg = avg / n

	for i := 0; i < n/2; i++ {
		result = result + (avg - ((2 * i) + 1))
	}

	return result
}
