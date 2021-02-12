package strings

import (
	"strings"
)

func stringToInt(s string)int{

	result,minus := 0,0
	var zeroRune rune
	zeroRune = '0'
	var minusRune rune  = '-' 

	s = strings.Trim(s, " ")
	for _,val := range s{

		num := val - zeroRune
		if int(num) > -1 && int(num) < 10{ 
		 result = result*10
		 result= result + int(num)
		 continue
		}
		if val == minusRune && result == 0{
           minus = -1
		}else{
			return 0
		}
	}

	if minus == -1{
		result = result * minus
	}

	return result
}