package assignment

import (
	"strconv"
)

// https://leetcode.com/problems/fizz-buzz/
func fizzbuzz(n int) []string {

    s := make([]string, n)
	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			s[i-1] =  "FIZZBUZZ"
		} else if i%3 == 0 {
			s[i-1] = "FIZZ"
		} else if i%5 == 0 {
			s[i-1] = "BUZZ"
		} else {
			s[i-1] = strconv.Itoa(i) 
		}

	}
	return s
}
