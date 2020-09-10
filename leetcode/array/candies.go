package array

func kidsWithCandies(candies []int, extraCandies int) []bool {

	length := len(candies)
	result := make ([]bool,length)
	maxPossible := 0

	for j, i := range candies {
		if i > maxPossible {
			maxPossible = i
		}
		candies[j] = i + extraCandies
	}

	for j, val := range candies {
		if val < maxPossible {
			result[j] = false
		} else {
			result[j] = true
		}
	}
	return result
}
