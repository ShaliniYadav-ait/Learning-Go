package leetcode

func diStringMatch(s string) []int {
	i:= 0
	min := 0
	max := 0
	len := len(s)
	max = len + 1
	result := make([]int, max)
	
    for  i=0 ;i<len; i++{
       if s[i] == 'I'{
		    result[i] = min
		    min = min + 1
	   }else{
		  
			result[i] = max - 1   
			max = max - 1
	   }
	}

	result[len] = max - 1
	return result
}