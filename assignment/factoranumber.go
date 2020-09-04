package assignment

func factorNumber(input []int, number int) []int {

	factor := []int{}

	for _, val := range input {
		for number%val == 0 && number != 0 {
            number = number / val
			factor = append(factor,val)
	}	
	}

	if number > 1{
		factor = append(factor,number)
	}
 return factor

}
