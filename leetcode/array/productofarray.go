package array

func productExceptSelf(nums []int) []int {

	mult := make([]int, len(nums))
	product := 1
	count := 0
	index := 0

	for i, val := range nums {
		if val == 0 {
			index = i
			count++
			if count > 1 {
				return mult
			}
			continue
		}

		product = product * val

	}

	if count == 1 {
		mult[index] = product
	}

	for i:=0; i<len(nums);i++{
		mult[i] = product/nums[i]
	}
	
	return mult
}

/* 	output := make([]int, len(nums))

	for _, i := range output {
		output[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			output[i] = output[i] * nums[j]
		}
	}
	return output
}
*/
