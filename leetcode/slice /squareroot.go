package slice

func findSquareRoot(input []int) (square []int) {

	i, j, p := 0, len(input)-1, len(input)-1
	square = make([]int, len(input))

	for i <= j {

		if input[i]*input[i] > input[j]*input[j] {
			square[p] = input[i] * input[i]
			i++
		} else {

			square[p] = input[j] * input[j]
			j--
		}
		p--
	}

	return square
}
