package array

func containsDuplicate(nums []int) bool {

	result := make(map[int]int)

	for _, i := range nums {
		val := result[i]
		if val+1 >1 {
			return true
		}
		result[i] = val + 1
	}
	return false

}
