package slice

func consecutiveOne(input []int) int {

	count := 0
	max := 0

	for _,val := range input {
		if val == 1 {
			count++    
        } else{
            count = 0
        }
        
		if max < count {
				max = count
		  }
	}
	return max 
}
