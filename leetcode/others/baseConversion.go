package others

import (
	"strconv"
	"unicode"
)

func baseConversion(s string, b1 int, b2 int) string {

	neg := false

	 num := 0

	if s[0] == '-' {
		neg = true
	}

	//for i := n ; i < len(s); i++{
	for i, val := range s {
		if i == 0 && neg == true {
			continue
		}

		num *= b1

		if unicode.IsDigit(val) {
			num = num + int(val) - '0'
		} else {
			num = num + int(val) - 'A' + 10
		}

	}

	return constructB2(num, b2)

}

func constructB2(num int, b int) string {
	if num == 0 {
		return "0"
	}

	var val string
	modVal := 0
	for num != 0{
	if (num % b) >= 10 {
		modVal = ('A' + (num % b) - 10)
		val = string(rune(modVal)) + val
	} else {
		modVal = num % b
		val = strconv.Itoa(modVal) + val
	}

	num = num / b
	}

	return val

}
