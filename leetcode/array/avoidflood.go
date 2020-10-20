package array

func avoidFlood(rains []int) []int {

	lakes := make(map[int]int)
	result := make([]int, len(rains))

	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			result[i] = 1
			continue
		}

		_, ok := lakes[rains[i]]
		if !ok {
			lakes[rains[i]] = i
			result[i] = -1
			continue
		}

		for j := lakes[rains[i]] + 1; j < i; j++ {
			if rains[j] == 0 {
				result[j] = rains[i]
				result[i] = -1
				rains[j] = -1
				lakes[rains[i]] = i
				break
			}
		}
		
		if lakes[rains[i]] != i {
			return []int{}
		}
	}

	return result
}
