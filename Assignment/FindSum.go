package assignment

func findSum(num []int) int{

	Sum := 0

	for _, i := range num {
		Sum = Sum + i
	}

	return Sum
}
