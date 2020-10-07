package assignment

// https://leetcode.com/problems/fibonacci-number/
func findFibonacci(num int) int {

	if num <= 1 {
		return num
	}

	i := 2
	num1 := 0
	num2 := 1
	fibo := 0

	for i <= num {
		fibo = num2 + num1
		num1 = num2
		num2 = fibo
		i++
	}

	return fibo
}
