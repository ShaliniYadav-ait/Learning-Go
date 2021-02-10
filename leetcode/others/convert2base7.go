package others

func convertToBase7(num int) string {
  	
    digits := []string{"0","1","2","3","4","5","6"}
	var result string 
	
    minus := false

    if num == 0{
        return "0"
    }
    
	if num < 0{
     minus = true
     num = num * -1
	}
    
    for num > 0{

		rem := num%7
		num = num/7

		result = digits[rem] + result 

	}

    if minus == true{
        result = "-" + result
    }
	return result  
}