package assignment

func anyBaseToAnyBase(fromBase int, toBase int, s string) string{

  var  fromDigits , toDigits []string	

  fromDigits , toDigits  = assignDigits(fromBase,toBase)

  number := anyBaseToDecimal(s, fromBase, fromDigits)
  result := decimalToAnyBase(number, toBase, toDigits)

return result

}

func assignDigits(fromBase int, toBase int)(from []string , to []string){

	switch fromBase{
	case 2 : from = []string{"0","1"}
	case 3 : from = []string{"0","1","2"}
	case 4 : from = []string{"0","1","2","3"}
	case 5 : from = []string{"0","1","2","3","4"}
	case 6 : from = []string{"0","1","2","3","4","5"}
	case 7 : from = []string{"0","1","2","3","4","5","6"}
	case 8 : from = []string{"0","1","2","3","4","5","6","7"}
	case 9 : from = []string{"0","1","2","3","4","5","6","7","8"}
	case 10 : from = []string{"0","1","2","3","4","5","6","7","8","9"}
	case 16 : from = []string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}
	default :     from = []string{"0","1","2","3","4","5","6","7","8","9"}
	} 

	switch toBase{
		case 2 : to = []string{"0","1"}
		case 3 : to = []string{"0","1","2"}
		case 4 : to = []string{"0","1","2","3"}
		case 5 : to = []string{"0","1","2","3","4"}
		case 6 : to = []string{"0","1","2","3","4","5"}
		case 7 : to = []string{"0","1","2","3","4","5","6"}
		case 8 : to = []string{"0","1","2","3","4","5","6","7"}
		case 9 : to = []string{"0","1","2","3","4","5","6","7","8"}
		case 10 : to = []string{"0","1","2","3","4","5","6","7","8","9"}
		case 16 : to = []string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}
		default :     to = []string{"0","1","2","3","4","5","6","7","8","9"}
		} 

		return from , to
}
