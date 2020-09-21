package assignment

func sortUsingIndexValue(nums []int, index int) []int {

	nums[index], nums[len(nums)-1] = nums[len(nums)-1], nums[index]

	k := 0
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] < nums[len(nums)-1] {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	nums[k] , nums[len(nums)-1] = nums[len(nums)-1] , nums[k]
	
	return nums
}
