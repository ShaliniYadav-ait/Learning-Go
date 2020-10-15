package array

// https://leetcode.com/explore/challenge/card/october-leetcoding-challenge/561/week-3-october-15th-october-21st/3496/
func rotate(nums []int, k int) []int{

	length := len(nums)
	for k != 0 {
		temp := nums[length-1]
		for i := length - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
		k--
	}
return nums
}
