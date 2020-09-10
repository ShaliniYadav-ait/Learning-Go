package array

func shuffle(nums []int, n int) []int {

	result := make([]int, len(nums))
	j := 1
	result[0] = nums[0]
	result[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums)/2; i++ {
		result[j] = nums[n+i-1]
		result[j+1] = nums[i]
		j = j + 2
	}
	return result
}
