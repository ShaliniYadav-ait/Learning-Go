package assignment

func findNumberWithGivenSum(input []int, sum int) []int {
	result := []int{-1, -1}
	firstIndex := 0

	for firstIndex < len(input) {

		tofind := sum - input[firstIndex]

		for index, val := range input {
			if val == tofind && index != firstIndex {
				result = []int{firstIndex, index}
				return result
			}

		}
		firstIndex++
	}

	return result
}
