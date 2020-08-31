package assignment

import (
	"strconv"
)

func fizzbuzz(n int) string {

	var s string
	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			s = s + "FIZZ BUZZ"
		} else if i%3 == 0 {
			s = s + "FIZZ"
		} else if i%5 == 0 {
			s = s + "BUZZ"
		} else {
			s = s + strconv.Itoa(i)
		}

		if i < n {
			s = s + ","
		}
	}
	return s
}
