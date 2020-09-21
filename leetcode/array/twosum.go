package array

func twoSum(nums []int, target int) []int {

	numsMap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		
		diff := target - nums[i]
		val, ok := numsMap[diff]
		
		if ok{
			return []int{i, val}
		}
		numsMap[nums[i]] = i
	}

	return nil
}
