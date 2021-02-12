package dailychallenge

func maxPower(s string) int {
	
	max,count := 0,0
    var prev string

	for _,val := range s{

		if prev == ""{
			prev = string(val)
			count = 1
			continue
		}

		if prev == string(val){
			count++
		}else{
			prev = string(val)
			if max < count{
				max = count
			}
			count = 1
		}
	}

	if max < count {
		max = count
	}

	return max
}