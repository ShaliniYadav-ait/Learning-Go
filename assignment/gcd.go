package assignment

func findGCD(num1 int, num2 int) int {

	if num2 > num1 {
		temp := num1
		num1 = num2
		num2 = temp
	}

	if num1%num2 == 0 {
		return num2
	}

return findGCD(num2, num1%num2)
}
