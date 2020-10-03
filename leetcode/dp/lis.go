package leetcode

func lengthOfLIS(nums []int) int {

	maxLen := make([]int, len(nums))
	globalMax := 0

	for i := 0; i < len(nums); i++ {
		localMax := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && localMax < maxLen[j]+1 {
				localMax = maxLen[j] + 1
			}
		}
		maxLen[i] = localMax
		if globalMax < localMax{
			globalMax = localMax
		}
	}

	return globalMax
}
