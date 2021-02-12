package dailychallenge

func longestMountain(A []int) int {

	longest, min, max, last := -1, -1, -1, -1

	for i := 0; i < len(A); i++ {

		if i == 0 {
			last = A[i]
			continue
		}

		if A[i] < last {
			if min == -1 {
				min = i
			} else {
				max = i
			}

			last = A[i]

			if longest < max-min{
				longest = max-min
			}
		}

		if A[i] > last {
			last = A[i]
			if max != -1{
				min = i - 1
			}
		}

	}

	return longest
}
