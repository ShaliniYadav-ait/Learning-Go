package leetcode

func validPalindrome(s string) bool {
	var pal bool
	len := len(s)
	totallen := len/2 + 1
len--

	for i := 0; i < (len / 2); i++ {
		if s[i] == s[totallen] && i <= (len/2) {
			pal = true
			totallen--
			continue
		} else {
			pal = false
		}
	}

	return pal
}
