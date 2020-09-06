package slice

func runningSumOfArray(input []int) (sum []int){

	total := 0
	for i := 0; i < len(input); i++ {
			total = total + input[i]
		    sum = append(sum, total)
	}
	return sum
}
