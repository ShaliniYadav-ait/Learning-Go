package slice

func findEvenDigits(input []int) int {

	digitCount, evenCount := 0,0

	for _, i := range input {
		digitCount = findDigits(i)

		if digitCount%2 == 0 {
			evenCount++
		}
	}
	return evenCount
}

func findDigits(num int) (digits int) {

	for num > 0 {
		num = num / 10
		digits++
	}

	return digits
}
